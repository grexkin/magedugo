package main

import (
	"fmt"
)

const cversion string = "1.2.3"

func main() {
	const funcVersion string = ""
	const pi = 3.1415926

	const (
		a1     string = "ddd"
		a2            = 122
		a3, a4 string = "ddd", "ddd"
	)

	const a5, a6 = "ddd", 9090
	// a5 = "dasads"
	fmt.Println(cversion, funcVersion, pi, a1, a2, a3, a4, a5, a6)
	const (
		e1 = "aaa"
		e2 //只定义标识符 使用上一行的常量标识符来赋值
		e3
		e4
		e5 = 1
		e6
		e7
	)
	fmt.Println(e1, e2, e3, e4, e5, e6, e7)

	const (
		f1 = iota
		f2
		f3
		f4
	)
	fmt.Println(f1, f2, f3, f4)
}
