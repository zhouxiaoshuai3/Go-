package fun_test

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T)  {
	defer Clear()
	fmt.Println("start")
	panic("fatal error")
}

func Clear()  {
	fmt.Println("clear resources")
}

//func sum(op ...int) int {
//	s := 0
//	for _, v := range op {
//		s += v
//	}
//	return s
//}
/**
可变参数
 */
//func TestVarParam(t *testing.T)  {
//	t.Log(sum(1,2,3,4)) // 10
//}