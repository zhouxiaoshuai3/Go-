package encap_test

import (
	"testing"
)

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {}
// 使用接口中 方法，方法签名一致，就实现了该接口
func (g *GoProgrammer) WriteHelloWorld()string  {
	return "fmt.Println(\"hello world\")"
}

func TestClient(t *testing.T)  {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld()) // fmt.Println("hello world")
}
