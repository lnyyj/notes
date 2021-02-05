package ttp

import (
	"fmt"
	"testing"
)


const (
   mutexLocked = 1 << iota
   mutexWoken
   mutexStarving
   mutexWaiterShift = iota   // mutexWaiterShift值为3，通过右移3位的位运算，可计算waiter个数
   starvationThresholdNs = 1e6 // 1ms，进入饥饿状态的等待时间
)

func Test_Lock1(t *testing.T) {
	fmt.Println("---->",mutexLocked)
	fmt.Println("---->",mutexWoken)
	fmt.Println("---->",mutexStarving)
	fmt.Println("---->",mutexWaiterShift)
	fmt.Println("---->",starvationThresholdNs)
}