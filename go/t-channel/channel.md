
#

channel跟string或slice有些不同，它在栈上只是一个指针，实际的数据都是由指针所指向的堆上面。

跟channel相关的操作有：初始化/读/写/关闭。channel未初始化值就是nil，未初始化的channel是不能使用的。下面是一些操作规则：

读或者写一个nil的channel的操作会永远阻塞。
读一个关闭的channel会立刻返回一个channel元素类型的零值。
写一个关闭的channel会导致panic。
map也是指针，实际数据在堆中，未初始化的值是nil。


