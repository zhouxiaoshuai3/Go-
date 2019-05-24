package loo_test

import (
	"fmt"
	"runtime"
	"testing"
)

func TestSwitch(t *testing.T)  {
	switch os := runtime.GOOS; os {
	case "darwin" :
		fmt.Println("OS X")
	case "linux":
		fmt.Printf("Linux")
	default:
		fmt.Printf("other ...")
	}
}

func TestSwitchMultiSe(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			fmt.Println(i, "偶数")
		case 1, 3:
			fmt.Println(i, "奇数")
		default:
			fmt.Println("it is not in 0~3")
		}
	}
}

func TestSwitchCondition(t *testing.T)  {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log(i, "偶数")
		case i%2 == 1:
			t.Log(i, "奇数")
		default:
			t.Log("unkwon")
		}
	}
}
/**
  switch_condition_test.go:37: 0 偶数
  switch_condition_test.go:39: 1 奇数
  switch_condition_test.go:37: 2 偶数
  switch_condition_test.go:39: 3 奇数
  switch_condition_test.go:37: 4 偶数
 */