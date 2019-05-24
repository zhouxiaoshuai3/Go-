package fib_test

import (
	"fmt"
	"testing"
)

var	(
	a , b = 1, 1
)

/**
	单元测试文件名 必须是以 _test.go 结尾。测试方法必须是 Test 关键字开始
	TestXXX 方法 接收参数是  t *testing T
 */


func TestFibList(t *testing.T) {
	fmt.Println(a)
	for i := 0;i < 5 ;i ++  {
		fmt.Println(b , " ")
		tmp := a
		a = b
		b = tmp + a
	}
	t.Log()
	fmt.Print()
}

func TestChangeVar(t *testing.T){
	a := 1
	b := 2
	a, b = 2, 1
	t.Log(a, b)
}
