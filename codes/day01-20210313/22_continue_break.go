package main

import "fmt"

func main() {
	//continue
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
	//break
	for j := 0; j < 10; j++ {
		if j == 3 {
			break
		}
		fmt.Println(j)
	}
}
