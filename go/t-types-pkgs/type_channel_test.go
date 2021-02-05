package ttp

import (
	"fmt"
	"testing"
	"time"
)

func Test_canlen(t *testing.T) {
	ch := make(chan int, 3)
	fmt.Println("ch ", len(ch), cap(ch))
	ch <- 1
	fmt.Println("ch ", len(ch), cap(ch))
	<-ch
	fmt.Println("ch ", len(ch), cap(ch))
}

func Test_Nil(t *testing.T) {
	var ch chan bool
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("------>make")
		ch = make(chan bool, 1)
	}()

	fmt.Println("------>write begin")
	ch <- true // 阻塞
	fmt.Println("------>write end")

	for {
		select {
		case v := <-ch:
			fmt.Println("--->", v)
		}
		break
	}
}
