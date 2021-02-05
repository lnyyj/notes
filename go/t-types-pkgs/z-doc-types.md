
# golang 常用类型解析

## channel

channel跟string或slice有些不同，它在栈上只是一个指针，实际的数据都是由指针所指向的堆上面。

跟channel相关的操作有：初始化/读/写/关闭。channel未初始化值就是nil，未初始化的channel是不能使用的。下面是一些操作规则：

读或者写一个nil的channel的操作会永远阻塞。
读一个关闭的channel会立刻返回一个channel元素类型的零值。
写一个关闭的channel会导致panic。
map也是指针，实际数据在堆中，未初始化的值是nil。


## slice
- slice 是一个栈上的指针,指向实际内容存放内存。 内存分配，大于32k分配在堆上，反之存放在P(GMP)的空闲内存链上

### 使用原则
1. 初始化时, 给一个合理的cap值。 因为随着slice不断的append, cap值会不断的突破内置预设，每次突破时，都会牵扯到内部数组拷贝。如: make([]int,0,100)
2. slice 的内部每个元素，也是存的每个元素地址。
3. 初始化深浅拷贝使用
    - 浅拷贝， 如下, 改动或者新增sli前五个元素都是会影响到arr

        ```
        var arr = [...]int{1, 2, 3, 4, 5}
        var sli = arr[:2]

        sli[0] = 10 // arr[0] 也会变更10    
        ```

    - 深拷贝， 改动或者变更sli不会影响arr

        ```
        var arr = [...]int{1, 2, 3, 4, 5}
        var sli = make([]int, 0, 4)
        sli = append(sli, arr[:2]...) // 深拷贝
        sli[0] = 10 // arr[0] 不变更
        ```



### 源码分析

- 内部定义, slice就是对array的操作

    ```
    type slice struct {
        array unsafe.Pointer
        len   int
        cap   int
    }
    ```

- 创建

    ```
    func makeslice(et *_type, len, cap int) unsafe.Pointer {
        // 计算分配内存大小
        mem, overflow := math.MulUintptr(et.size, uintptr(cap))
        // 是否溢出判断
        if overflow || mem > maxAlloc || len < 0 || len > cap {
            mem, overflow := math.MulUintptr(et.size, uintptr(len))
            if overflow || mem > maxAlloc || len < 0 {
                panicmakeslicelen()
            }
            panicmakeslicecap()
        }
        
        // 根据需要内存大小, 确定内存分配位置
        // mem >  32k 分配在堆上上
        // mem <= 32k 分配在P的空闲内存上(多级向上分配: pre-p ---> m(忘记了) ----> global ---> heap)
        return mallocgc(mem, et, true)
    }
    ```

- 新增

```
func growslice(et *_type, old slice, cap int) slice {
	if raceenabled {
		callerpc := getcallerpc()
		racereadrangepc(old.array, uintptr(old.len*int(et.size)), callerpc, funcPC(growslice))
	}
	if msanenabled {
		msanread(old.array, uintptr(old.len*int(et.size)))
	}

	if cap < old.cap {
		panic(errorString("growslice: cap out of range"))
	}

    // 如果存储的类型空间为0，  比如说 []struct{}, 数据为空
	if et.size == 0 {
		return slice{unsafe.Pointer(&zerobase), old.len, cap}
	}

	newcap := old.cap
	doublecap := newcap + newcap

    // 计算扩容后的cap大小
	if cap > doublecap { // 如果新的cap大于原来的cap两杯
		newcap = cap
	} else {
		if old.len < 1024 { // len小于1024, 成2倍增加
			newcap = doublecap
		} else {
            // 大于1024， 成1.25倍增加
			for 0 < newcap && newcap < cap {
				newcap += newcap / 4
			}
			if newcap <= 0 { // 溢出判断
				newcap = cap
			}
		}
	}


    // 对不同的类型，计算分配内存大小

	var overflow bool
	var lenmem, newlenmem, capmem uintptr
	// Specialize for common values of et.size.
	// For 1 we don't need any division/multiplication.
	// For sys.PtrSize, compiler will optimize division/multiplication into a shift by a constant.
	// For powers of 2, use a variable shift.
	switch {
	case et.size == 1:
		lenmem = uintptr(old.len)
		newlenmem = uintptr(cap)
		capmem = roundupsize(uintptr(newcap))
		overflow = uintptr(newcap) > maxAlloc
		newcap = int(capmem)
	case et.size == sys.PtrSize:
		lenmem = uintptr(old.len) * sys.PtrSize
		newlenmem = uintptr(cap) * sys.PtrSize
		capmem = roundupsize(uintptr(newcap) * sys.PtrSize)
		overflow = uintptr(newcap) > maxAlloc/sys.PtrSize
		newcap = int(capmem / sys.PtrSize)
	case isPowerOfTwo(et.size):
		var shift uintptr
		if sys.PtrSize == 8 {
			// Mask shift for better code generation.
			shift = uintptr(sys.Ctz64(uint64(et.size))) & 63
		} else {
			shift = uintptr(sys.Ctz32(uint32(et.size))) & 31
		}
		lenmem = uintptr(old.len) << shift
		newlenmem = uintptr(cap) << shift
		capmem = roundupsize(uintptr(newcap) << shift)
		overflow = uintptr(newcap) > (maxAlloc >> shift)
		newcap = int(capmem >> shift)
	default:
		lenmem = uintptr(old.len) * et.size
		newlenmem = uintptr(cap) * et.size
		capmem, overflow = math.MulUintptr(et.size, uintptr(newcap))
		capmem = roundupsize(capmem)
		newcap = int(capmem / et.size)
	}

	// The check of overflow in addition to capmem > maxAlloc is needed
	// to prevent an overflow which can be used to trigger a segfault
	// on 32bit architectures with this example program:
	//
	// type T [1<<27 + 1]int64
	//
	// var d T
	// var s []T
	//
	// func main() {
	//   s = append(s, d, d, d, d)
	//   print(len(s), "\n")
	// }
	if overflow || capmem > maxAlloc {
		panic(errorString("growslice: cap out of range"))
	}

    // 内存分配
	var p unsafe.Pointer
	if et.ptrdata == 0 {
		p = mallocgc(capmem, nil, false)
		// The append() that calls growslice is going to overwrite from old.len to cap (which will be the new length).
		// Only clear the part that will not be overwritten.
		memclrNoHeapPointers(add(p, newlenmem), capmem-newlenmem)
	} else {
		// Note: can't use rawmem (which avoids zeroing of memory), because then GC can scan uninitialized memory.
		p = mallocgc(capmem, et, true)
		if lenmem > 0 && writeBarrier.enabled {
			// Only shade the pointers in old.array since we know the destination slice p
			// only contains nil pointers because it has been cleared during alloc.
			bulkBarrierPreWriteSrcOnly(uintptr(p), uintptr(old.array), lenmem)
		}
	}

    // 数据拷贝
	memmove(p, old.array, lenmem)

	return slice{p, old.len, newcap}
}
```

