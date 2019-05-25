package string_test

import (
	"strconv"
	"testing"
)

func TestConv(t *testing.T)  {
	//s := strconv.Itoa("10") // int 转换 字符串
	if val, ok := strconv.Atoi("10"); ok == nil {
		t.Log(val + 10)
	} else {
		t.Log("fail")
	}
	//t.Log(s + "str") // 10str
}

//func TestStringFun(t *testing.T)  {
//	s := "a, b, c, d"
//	s_slice := strings.Split(s, ",") // string 转 slice
//	t.Log(s_slice) // [a  b  c  d]
//	t.Log(strings.Join(s_slice, "-")) // 字符串连接  join  a- b- c- d
//}
//
//
//
//func TestString(t *testing.T)  {
//	var s string
//	//t.Log(s) // 初始化为默认零值 ""
//	//s = "hello"
//	//t.Log(len(s)) // byte 数是 5
//	//s[1] = "s" // string 是不可变的 byte slice
//	s = "\xE4\xB8\xA5" // 可存储任何二进制数据
//	t.Log(s) // 严
//	//t.Log(len(s)) // len 是 byte 数  3，并不一定是字符数
//	s = "中"
//	t.Log(len(s)) // 是 byte 数  3
//	c := []rune(s)
//	t.Log(len(c)) // rune的 unicode 编码 1
//	t.Logf("中 unicode %x", c[0]) // %x 是输出 十六进制 4e2d
//	t.Logf("中 UTF8 %x", s) // utf8 的输出  e4b8ad
//}
//
//func TestStringToRune(t *testing.T)  {
//	s := "中华人名共和国"
//	for _, c := range s{
//		t.Logf("%[1]c %[1]d", c) // 占位符 [1] 的意思是 两个占位符都是和第一个参数对应
//	}
//	/**
//		string_test.go:25: 中 20013
//	    string_test.go:25: 华 21326
//	    string_test.go:25: 人 20154
//	    string_test.go:25: 名 21517
//	    string_test.go:25: 共 20849
//	    string_test.go:25: 和 21644
//	    string_test.go:25: 国 22269
//	 */
//}
//
//func TestConv(t *testing.T)  {
//	s := strconv.Itoa(10) // int 转换 字符串
//	t.Log(s + "str") // 10str
//}