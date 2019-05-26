package encap_test

import (
	"fmt"
	"testing"
)

type Employee struct {
	Id string
	Name string
	Age int
}

func (e Employee) String()string  {
	return fmt.Sprintf("id:%s-name:%s-age:%d", e.Id, e.Name, e.Age)
}
func (e *Employee) Strings()string  {
	return fmt.Sprintf("id:%s-name:%s-age:%d", e.Id, e.Name, e.Age)
}

func TestStructOp(t *testing.T)  {
	e := Employee{"0", "Bob", 20}
	t.Log(e.String()) // id:0-name:Bob-age:20
	t.Log(e.Strings()) // 实例得到的是一个类型的指针 id:0-name:Bob-age:20
}

//func TestCreateEmployee(t *testing.T) {
//	e := Employee{"0", "Bob", 20}
//	e1 := Employee{Name:"Mike", Age:19}
//	e2 := new(Employee) // 创建的 结构体是指针类型
//	e2.Id = "2"
//	e2.Name = "Rose"
//	e2.Age = 30
//	t.Log(e)
//	t.Log(e1)
//	t.Log(e1.Id)
//	t.Log(e2)
//	t.Logf("e is %T", e) // e is encap_test.Employee
//	t.Logf("e is %T", &e) // e is *encap_test.Employee 加上 & 符，就是指针类型
//	t.Logf("e2 is %T", e2) // %T 输出类型 e2 is *encap_test.Employee
//}