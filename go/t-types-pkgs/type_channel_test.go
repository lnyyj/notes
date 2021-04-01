package ttp

import (
	"fmt"
	"testing"
	"time"
)

func Test_in_out(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch2 <- 1
	}()

	go func() {
		fmt.Println("11111")
		ch1 <- <-ch2
		fmt.Println("33333")
	}()
	go func() {
		ch1 <- <-ch2
		fmt.Println("2222")
	}()
	<-ch1
}

func Test_isClose(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	_, ok := <-ch
	fmt.Println("------>ok: ", ok)
	close(ch)
	v, ok := <-ch
	fmt.Printf("------>v: [%+v][%+v]\r\n", ok, v)
}

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
