package main

import "fmt"

func main() {
	var confirmf string = "y"
	var i int = 0
	for confirmf == "y" {
		fmt.Println(i)
		i++
		fmt.Print("继续吗？（y/n）")
		fmt.Scan(&confirmf)
	}
	fmt.Println("over")
}
