package main

import "fmt"

func main() {
	//简短声明，通过值来推导标识符类型
	name := "jsonKk"
	fmt.Println(name)
	//整数： int，byte
	id := 1
	fmt.Println(id)
	//交换值
	a1, a2 := 1, 2
	fmt.Println(a1, a2)
	a1, a2 = a2, a1
	fmt.Println(a1, a2)
}
