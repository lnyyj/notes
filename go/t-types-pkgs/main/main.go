package main

import (
	"fmt"
	"runtime"
	"time"
)

func f() {
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

func f1() {
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

// type People struct{}

// func (p *People) ShowA() {
// 	fmt.Println("showA")
// 	p.ShowB()
// }
// func (p *People) ShowB() {
// 	fmt.Println("showB")
// }

// type Teacher struct {
// 	People
// }

// // func (t *Teacher) ShowA() {
// // 	fmt.Println("teacher showA")
// // 	t.ShowB()
// // }
// func (t *Teacher) ShowB() {
// 	fmt.Println("teacher showB")
// }
// func f3() {
// 	t := Teacher{}
// 	t.ShowA()
// }

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func f4() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

type People interface {
	Speak(string) string
}
type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}
func f5() {
	var peo People = &Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

func f6() {
	ch1 := make(chan struct{})
	go func() {
		for i := 1; i < 100; i = i + 2 {
			// if i%2 == 0 {
			fmt.Println("A:", i)
			// }

			ch1 <- struct{}{}
		}
	}()
	go func() {
		// for i := 1; i < 100; i++ {
		for i := 2; i < 100; i = i + 2 {
			runtime.Gosched()
			// if i%2 == 1 {
			fmt.Println("B:", i)
			// }
			<-ch1
		}
	}()
	time.Sleep(1 * time.Second)
	close(ch1)
}

func maxEnvelopes(envelopes [][]int) int {

	gt := func(src []int, dst []int, eq bool) (ok bool) {
		if len(src) != len(dst) {
			return false
		}
		ok = true
		for i := 0; i < len(src); i++ {
			if eq {
				if src[i] < dst[i] {
					return false
				}
			} else {
				if src[i] <= dst[i] {
					return false
				}
			}
		}
		return
	}

	// 找到最大信封
	var maxinx int
	l := len(envelopes)
	if l == 0 {
		return 0
	}
	for i := 1; i < l; i++ {
		if ok := gt(envelopes[0], envelopes[i], true); !ok {
			maxinx = i
		}
	}

	var tmps [][]int
	tmps = append(tmps, envelopes[maxinx])
	for inx, v := range envelopes {
		if inx != maxinx && gt(envelopes[maxinx], v, false) {
			l := len(tmps)
			if l == 1 {
				tmps = append(tmps, v)
			} else {
				for tmpInx := 1; tmpInx < l; tmpInx++ {
					if gt(v, tmps[tmpInx], false) {
						tmps = append(tmps[:tmpInx], append([][]int{v}, tmps[tmpInx:]...)...)
					} else if tmpInx == l-1 {
						tmps = append(tmps, v)
					}
				}
			}
		}
	}
	return len(tmps)
}

func main() {
	// f1()
	// f3()
	// f4()
	// f6()

	envelopes := make([][]int, 0, 4)
	// envelopes = append(envelopes, []int{5, 4}, []int{6, 4}, []int{6, 7}, []int{2, 3})
	envelopes = append(envelopes, []int{4, 5}, []int{4, 6}, []int{6, 7}, []int{2, 3}, []int{1, 1})

	fmt.Println("------->", maxEnvelopes(envelopes))
}
