package fun_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)
/**
函数执行时间函数
*/
func timeSpent(inner func(op int)int) func(op int) int  {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent :", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFun(t *testing.T)  {
	// 函数多个返回值，使用 a, b接收.可以使用 _ 忽略 单个返回值
	//a, b := returnMultiValue()
	//_, d := returnMultiValue()
	//t.Log(a, b, d) // 1 7 19
	tsFun := timeSpent(slowFun)
	t.Log(tsFun(10))
}

func returnMultiValue() (int, int)  {
	return rand.Intn(10) , rand.Intn(20)
}



