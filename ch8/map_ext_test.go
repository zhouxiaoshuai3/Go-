package map__test

import "testing"

func TestMapWithFunValue(t *testing.T)  {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}
	// 调用 m[1] 的value【function】并传递参数 2
	t.Log(m[1](2), m[2](2), m[3](2)) // 2 4 8
}
