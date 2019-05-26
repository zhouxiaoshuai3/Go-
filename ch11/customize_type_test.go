package encap_test

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv  {
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Println("time spent :", time.Since(start).Seconds())
		return ret
	}
}
func slowFun(op int)int  {
	time.Sleep(time.Second * 1)
	return op
}

func TestFun(t *testing.T)  {
	sfFun := timeSpent(slowFun)
	t.Log(sfFun(10))
}
