package type_test

import "testing"

/**
	1。 各种类型的隐式类型转换不支持，别名也不支持隐式转换
 */
func TestType(t *testing.T)  {
	var a int32 = 1
	var b int64 = 2
	b = int64(a)
	t.Log(a, b)
}
/**
	2。 指针类型，指针类型不支持运算
	3。 字符串类型，是值类型，字符串的初始值为空字符串
 */