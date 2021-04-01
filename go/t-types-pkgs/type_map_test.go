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

func Test_ChangeMap(t *testing.T) {
	{
		ageMp := make(map[string]int)
		ageMp["qcrao"] = 18
		for name := range ageMp {
			delete(ageMp, name)
			ageMp[name+name] = 1
			fmt.Println("---> ", name, len(ageMp))
		}
		fmt.Println(ageMp)
	}
	fmt.Println("==============================================")
	{
		ageMp := make(map[string]int)
		ageMp["qcrao"] = 18
		for name := range ageMp {
			ageMp[name+name] = 1
			delete(ageMp, name)
			fmt.Println("---> ", name, len(ageMp))
		}
		fmt.Println(ageMp)
	}

}
