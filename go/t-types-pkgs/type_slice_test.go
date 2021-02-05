package ttp

import (
	"fmt"
	"sync"
	"reflect"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func Test_Append(t *testing.T) {
	sli := make([]struct{}, 0 , 10000)
	ch := make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(100)
	for i:=0; i<100; i++ {
		go func(){
			<-ch
			for j:=0;j<100;j++{
				sli = append(sli, struct{}{})
			}
			fmt.Println("----->start: ", i)
			wg.Done()
		}()
	}
	close(ch)
	wg.Wait()
	fmt.Println("----->", len(sli))

	select{}
}


func Test_Range(t *testing.T) {
	var arr = [3]*User{&User{"ssy", 18}, &User{"ltt", 17}, &User{"syx", 2}}
	sli := make([]*User, 0, 3)
	for _, item := range arr {
		sli = append(sli, item)
	}
	for i:=0; i<3; i++ {
		fmt.Println("----->",sli[i])
	}

}

func Test_Struct(t *testing.T) {
	var arr = [...]User{User{"ssy", 18}, User{"ltt", 17}, User{"syx", 2}}
	// var sli = arr[:2]  // 浅拷贝
	sli := make([]User, 0)
	sli = append(sli, arr[:2]...) // 深拷贝

	show := func() {
		fmt.Printf("---->arr: %+v %+v %+v %+v %p %p \r\n", reflect.TypeOf(arr).Kind(), len(arr), cap(arr), arr, &arr, &arr[0])
		fmt.Printf("---->sli: %+v %+v %+v %+v %p %p \r\n", reflect.TypeOf(sli).Kind(), len(sli), cap(sli), sli, &sli, &sli[0])
	}

	show()
	sli[0].Age = 10
	show()
	sli = append(sli, User{"sxh", 48})
	show()
	sli = append(sli, User{"stp", 49})
	show()
}

func Test_append(t *testing.T) {
	var arr = []int{1, 2, 3, 4, 5}
	var sli = arr[:2]

	show := func() {
		fmt.Printf("---->arr: %+v %+v %p %p \r\n", reflect.TypeOf(arr).Kind(), arr, &arr, &arr[0])
		fmt.Printf("---->sli: %+v %+v %p %p \r\n", reflect.TypeOf(sli).Kind(), sli, &sli, &sli[0])
	}
	show()
	sli[0] = 10
	show()
	sli = append(sli, 6, 7)
	show()
	sli = append(sli, 8, 9)
	show()
}

func Test_param(t *testing.T) {
	var sli = []int{1, 2, 3, 4, 5}

	input := func(v []int) {
		fmt.Printf("---->sli: %+v %+v %+v %p %p \r\n", len(v), cap(v), v, &v, &v[0])
		v = append(v, 6)
	}
	show := func() {
		fmt.Printf("---->sli: %+v %+v %+v %p %p \r\n", len(sli), cap(sli), sli, &sli, &sli[0])
	}
	show()
	input(sli)
	show()
}

func Test_Init(t *testing.T) {
	// sli := make([]int, 0, 5)
	sli := make([]int, 0)

	show := func() {
		fmt.Printf("---->sli: %+v %+v %+v %p %p \r\n", len(sli), cap(sli), sli, &sli, &sli[0])
	}

	// show()
	for i := 1; i <= 8; i++ {
		sli = append(sli, i)
		show()
	}

}
