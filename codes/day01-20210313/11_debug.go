package main

import "fmt"

func main() {
	fmt.Println("a", "b")
	fmt.Println("1", "2")
	fmt.Print("a", "b")
	fmt.Print("1", "2")
	//格式化字符串
	//占位符
	/*
		%T,%v,%#v与数据类型无关
		%v,%#v以可识别的格式打印
		%s  字符串
		%d  整数
	*/
	var name string = "ddd"
	var id = 1
	fmt.Printf("name:%T,ID:%T\n", name, id)
	fmt.Printf("name:%v,ID:%v\n", name, id)
	fmt.Printf("name:%#v,ID:%#v\n", name, id)
	fmt.Printf("name:%#s,ID:%#d\n", name, id)
}
