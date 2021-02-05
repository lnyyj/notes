package ttp

import (
	"fmt"
	"reflect"
	"testing"
)

func show1(u interface{}) {
	switch u.(type) {
	case *User:
		uu := reflect.ValueOf(u).Interface().(*User)
		fmt.Printf("--->func[%p] [%p] [%+v]\r\n", u, uu, u) // u 的地址变了。 但是uu地址没有变变化，指向了外部u的地址
	default:
		uu := reflect.ValueOf(u).Interface().(User)
		fmt.Printf("--->func[%p] [%p] [%+v]\r\n", &u, &uu, u) // u = uu != 外部u
	}
}

func show2(u User) {
	fmt.Printf("--->func[%p] [%+v]\r\n", &u, u) // u 的地址已变化
}
func show3(u *User) {
	fmt.Printf("--->func[%p] [%+v]\r\n", u, u) // 地址无变化
}

func Test_Param(t *testing.T) {
	u := User{"name", 20}
	fmt.Printf("---> [%p][%+v]\r\n", &u, u)
	show1(u)
	show1(&u)
	show2(u)
	show3(&u)
}

func b_show1(u interface{}) {

}

func b_show2(u User) {
	u.Name = "ssy"
}
func b_show3(u *User) {
	u.Name = "ssy"
}

func Benchmark_interfaceParam(b *testing.B) {
	// Benchmark_interfaceParam-4   	1000000000	         0.483 ns/op	       0 B/op	       0 allocs/op
	// u := User{"name", 20}
	// Benchmark_interfaceParam-4   	1000000000	         0.472 ns/op	       0 B/op	       0 allocs/op
	u := &User{"name", 20}
	for i := 0; i < b.N; i++ {
		b_show1(u)
	}
}

func Benchmark_sourceTypeParam(b *testing.B) {
	u := User{"name", 20}
	// Benchmark_sourceTypeParam-4   	1000000000	         0.315 ns/op	       0 B/op	       0 allocs/op
	for i := 0; i < b.N; i++ {
		b_show2(u)
	}
}
func Benchmark_sourceTypeAddrParam(b *testing.B) {
	u := &User{"name", 20}
	// Benchmark_sourceTypeAddrParam-4   	1000000000	         0.295 ns/op	       0 B/op	       0 allocs/op
	for i := 0; i < b.N; i++ {
		b_show3(u)
	}
}
