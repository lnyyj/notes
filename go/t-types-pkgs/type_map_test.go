package ttp

import (
	"fmt"
	"testing"
)

func Test_RangeMap(t *testing.T) {
	a := make(map[string]string, 0)
	a["aa"] = "aa1"
	a["bb"] = "bb1"
	for _, v := range a {
		fmt.Println("------>", v)
	}
}
