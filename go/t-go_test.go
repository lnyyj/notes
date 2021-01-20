package tgo

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test_Slice(t *testing.T) {
	s1 := make([]string, 0)
	s1 = append(s1, "1111")
	s1 = append(s1, "2222")
	s1 = append(s1, "3333")
	s1 = append(s1, "4444")
	s1 = append(s1, "5555")

	type xxxx []interface{}
	ssss := (*xxxx)(unsafe.Pointer(&s1))

	for _, v := range *ssss {
		fmt.Println("----->s:", v)
	}

	// fmt.Printf("----->s1:[%+v]\r\n", s1)
	// s2 := s1
	// fmt.Printf("----->s2:[%+v]\r\n", s2)
	// s1 = nil
	// fmt.Printf("----->s1:[%+v]\r\n", s1)
	// fmt.Printf("----->s2:[%+v]\r\n", s2)
}

func Test_Print(t *testing.T) {
	var v interface{}
	var a float64 = 1234567891234567891
	v = a

	fmt.Printf("------->[%19f][%v][%+v]", v, uint64(v.(float64)), v)
}
