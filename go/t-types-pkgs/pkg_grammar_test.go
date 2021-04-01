package ttp

import (
	"fmt"
	"testing"
)

type IT interface {
	ShowA()
	ShowB()
}

func Test_switch(t *testing.T) {
	a := 2
	switch a {
	case 1:
		println("11111")
		fallthrough
	case 2:
		println("22222")
		fallthrough
	case 3:
		println("333333")
		// fallthrough
	default:
		println("22222")
	}
}

func Test_DeferPanic(t *testing.T) {
	defer func() {
		fmt.Println("1")
	}()
	defer func() {
		fmt.Println("2")
		panic("2.222222")
	}()
	defer func() {
		fmt.Println("3")
		panic("3.3333333")
	}()
	panic("444")
	fmt.Println("5")
}

func Test_DeferRecover(t *testing.T) {
	defer func() {
		fmt.Println("1")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	defer func() {
		fmt.Println("3")
		panic("3.3333333")
	}()
	panic("444")
}
