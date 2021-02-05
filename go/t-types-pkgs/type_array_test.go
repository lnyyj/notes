package ttp

import (
	"fmt"
	"testing"
)

func Test_ArrayInit(t *testing.T) {
	a1 := [4]int{1, 2, 3, 4}   // 初始化时，直接分配栈内存
	a2 := [...]int{1, 2, 3, 4} // 这种方式不好，需要内部进行上限推导
	a3 := [5]int{1, 2, 3, 4}   // 先在静态存储区初始化，然后拷贝在栈上
	fmt.Println("a", a1, a2, a3)
}
