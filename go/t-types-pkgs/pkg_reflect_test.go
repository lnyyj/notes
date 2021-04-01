package ttp

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_law1(t *testing.T) {

}

type ITT interface {
	DO()
}

type B struct {
}

func (b B) DO() {
	fmt.Println("---->do")
}

func Test_Implements(t *testing.T) {
	_typer := reflect.TypeOf((*ITT)(nil)).Elem()
	var b B
	rv := reflect.ValueOf(&b)
	ok := rv.Type().Implements(_typer)
	fmt.Println("---->", ok)
	rv.MethodByName("DO").Call([]reflect.Value{})

}
