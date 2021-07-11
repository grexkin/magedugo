package main

import "fmt"

func main() {
	var (
		f1 float64 = 3.1
		f2 float64 = 3.2
		f3 float64 = 1e-5
	)
	fmt.Printf("%T,%v\n", f1, f1)
	fmt.Printf("%T,%v\n", f2, f2)
	fmt.Printf("%T,%v\n", f3, f3)

	fmt.Printf("%f,%e\n", f1, f1)
}
