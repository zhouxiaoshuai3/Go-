package main

import (
	"fmt"
	"os"
)

// hello world
// 接收命令行 参数  os.Args[]
// main 函数的返回值 os.Exit
func main() {
	arg := os.Args
	//exis := os.Exit(0)
	fmt.Println(arg)
}
