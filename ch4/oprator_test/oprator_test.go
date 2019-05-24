package oprator_test

import "testing"

/**
数组比较时，比较的数组必须长度，类型，和值都相等才是 true，否则都是 false
长度不同的两个数组，不能比较，编译时就回出错。原因是两个长度不同的数组是不同的类型
*/
func TestOpratorArr(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	b := [...]int{1, 2, 3, 5, 3}
	//c := [...]int{1,2,3,4}
	//d := [...]int{1,2}
	e := [...]int{1, 2, 3, 4, 5}
	t.Log(a == b)
	t.Log(a == e)
	//t.Log(a == c)
	//t.Log(a == d)
}

const (
	kedu = 1 >> iota
	kexie
	kezhixing
)

/**
&^ 比较符，按位清零
当比较符右边的二进制位上的值是 1，结果都会变成0。如果右边二进制位上的值是 0，结果以左边的值为准。
*/
func TestByteClear(t *testing.T) {
	a := 7 //
	a = a &^ kedu
	t.Log(a&kedu)
	t.Log(a&kexie)
	t.Log(a&kezhixing)
}
