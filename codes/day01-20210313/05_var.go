package main

import "fmt"

//显式赋值
var version string = "1.1.1"

func main() {
	var funcVersion string
	fmt.Println(funcVersion, version)
	var name = "ddd"
	var (
		a1     int
		a2     string
		a3            = 4
		a4     string = "ddd"
		a5, a6        = "ddd", 123
	)
	fmt.Println(name)
	fmt.Println(a1, a2, a3, a4, a5, a6)
	//自动推导
	var a7, a8 string = "a7", "a8"
	fmt.Println(a7, a8)
}
