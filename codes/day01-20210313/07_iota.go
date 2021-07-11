package main

import "fmt"

func main() {
	const (
		a1 = iota
		a2 = 9
		a3
		a4 = iota
		a5
		a6
	)
	fmt.Println(a1, a2, a3, a4, a5, a6)
}
