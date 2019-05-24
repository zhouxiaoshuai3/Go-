package constant_test

import "testing"

const (
	zhouyi = iota + 1
	zhouer
	zhousan
)
const (
	kedu = 1 >> iota
	kexie
	kezhixing
)
func TestConstantTest(t *testing.T) {
	t.Log(zhouyi, zhouer, zhousan)
}
func TestReadAnd(t *testing.T)  {
	a := 7
	t.Log(a&kedu, a&kexie, a&kezhixing)

}
