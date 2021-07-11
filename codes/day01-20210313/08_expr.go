package main

import "fmt"

func main() {
	var a1 = 10
	var a2 = a1 + 10
	var a3 = a1 * a2
	fmt.Println(a3)

	const e1 = 10
	const e2 = e1 + 10
	const e3 = e1 * e2
	fmt.Println(e3)

	var a4 = a3 + e3
	fmt.Println(a4)
}
