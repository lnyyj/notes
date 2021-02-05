# golang 内置包解析

## 调度器
- mpg, m 表示线程, p 处理器(调度的中心)， g协程(内存消耗2K左右)

### 调度方式
#### 循环调度
#### 阻塞调度
#### 抢占调度

### goroution 创建
- 获取函数 newProc ---> newProc1 ---> gfget ---> new(2k)
    - 首先使用gfget获取一个已经释放的g(gfget 现在当前p获取(32～64个)， 获取不到就去全局获取. p没有，全局有，就拿32个过来。)
    - 如果获取不到，就new一个
### 放入runq
    - runqput 
        pre-p 最有只能存放256个g, 多余的放在全局队列里
        
### goroution 回收
    - 使用fgput函数（本地g超过64个，在把多余的拷贝到全局去）

## reflect 
- 三大定律
    1. 从interface类型 到 reflect.Type/reflect.Value的转变
    2. 从reflect.Value到interface的转变
    3. 要修改反射对象，其值必须可设置。 即：可以通过reflect.Value变更其内部的元素值

### 使用要点
    1. 会用一定的性能损耗
    2. 运行时触发，编译时无法检错
    3. 总体结构是树型结构

具体介绍: https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-reflect/
    


## 锁

- 悲观锁/乐观锁是指锁的用法，是一个概念，并没有锁的实体。
    - 悲观锁-每次操作数据之前都加锁
    - 乐观锁-每次操作的时候不加锁，而是通过一个ver来判断是否操作数据

### 互斥锁 sync.Mutex

- 结构定义如下， 原子性操作state位进行加锁解锁. 有两种模式
    - 正常模式(可争抢) 
        1. waiter按照先进先出（FIFO）的方式获取锁
        2. 这个模式下会导致锁的争抢，新建gorouting和刚被唤醒gorouting对一个锁争抢，因为新建立的gorouting持有CP，所有后者基本上抢不赢，这个时候就会将锁切换成饥饿模式。
        3. 可以自旋等待锁，自旋条件(多cpu,自旋次数小于4,可以马上持有p(p持有M,切p的runq为空，当前goroution可以立刻切入) 等等)
        4. 自旋时会空转CPU 
    - 饥饿模式(排队拿锁)
        1. 锁的所有权直接从解锁（unlocking）的goroutine转移到等待队列中的队头waiter。新来的goroutine不会尝试去获取锁，也不会自旋。它们将在等待队列的队尾排队。
        2. 根据waiter的值，来确定队列位置
        3. 使用sema信号量来控制
        4. 拿锁时间小于1ms就切换成正常模式
        5. 队列中最后一个waiter持有了锁，也切换为正常模式

```
type Mutex struct {
	// 低三位的对应， locked, woken, starving状态
	// 高29位，表示waiter的数量
	state int32
	sema  uint32 // 用于控制锁状态的信号量。
}
```

### 读写锁 - sync.RWMutex

```
type RWMutex struct {
	w           Mutex  // held if there are pending writers
	writerSem   uint32 // 写信号 semaphore for writers to wait for completing readers
	readerSem   uint32 // 读信号 semaphore for readers to wait for completing writers
	readerCount int32  // 存储了当前正在执行的读操作数量 number of pending readers
	readerWait  int32  // 表示当写操作被阻塞时等待的读操作个数 number of departing readers
}

```

虽然读写互斥锁 sync.RWMutex 提供的功能比较复杂，但是因为它建立在 sync.Mutex 上，所以实现会简单很多。我们总结一下读锁和写锁的关系：

- 调用 sync.RWMutex.Lock 尝试获取写锁时；
    - 每次 sync.RWMutex.RUnlock 都会将 readerCount 其减一，当它归零时该 Goroutine 会获得写锁；
    - 将 readerCount 减少 rwmutexMaxReaders 个数以阻塞后续的读操作；
- 调用 sync.RWMutex.Unlock 释放写锁时，
    - 会先通知所有的读操作，然后才会释放持有的互斥锁；
    
- 调用 sync.RWMutex.RLock时，
    - 其实就是在读计数（readerCount）上+1， 如果+1后小于0，就证明加了写锁，就需要等待监听读信号量

- 调用 sync.RWMutex.RUnlock时
    - 其实就是在读计数（readerCount）上-1， 如果-1后小于0，就证明与写锁等待，就需要通知写信号量
    
读写互斥锁在互斥锁之上提供了额外的更细粒度的控制，能够在读操作远远多于写操作时提升性能。


### sync.WaitGroup    
- 对 sync.WaitGroup 的分析和研究，我们能够得出以下结论：
    - sync.WaitGroup 必须在 sync.WaitGroup.Wait 方法返回之后才能被重新使用；
    - sync.WaitGroup.Done 只是对 sync.WaitGroup.Add 方法的简单封装，我们可以向 sync.WaitGroup.Add 方法传入任意负数（需要保证计数器非负）快速将计数器归零以唤醒等待的 Goroutine；
    - 可以同时有多个 Goroutine 等待当前 sync.WaitGroup 计数器的归零，这些 Goroutine 会被同时唤醒；