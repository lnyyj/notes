package tchannel

import (
	"fmt"
	"testing"
)

var ch chan bool

func Test_Nil(t *testing.T) {
	fmt.Println("------>write begin")
	ch <- true
	fmt.Println("------>write end")

	for {
		select {
		case v := <-ch:
			fmt.Println("--->", v)
		}
		break
	}
}
