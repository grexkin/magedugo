package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("break")
BEXIT:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				break BEXIT
			}
			fmt.Println(i, j)
		}
	}
	fmt.Println(strings.Repeat("=", 30))
CEXIT:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				continue CEXIT
			}
			fmt.Println(i, j)
		}
	}
}
