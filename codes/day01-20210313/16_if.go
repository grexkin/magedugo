package main

import "fmt"

func main() {
	var confirm string
	fmt.Println("买10个包子")

	fmt.Print("有卖西瓜的吗？（y/n）")
	fmt.Scan(&confirm)

	if confirm == "y" {
		fmt.Println("买1个西瓜")
	} else {
		fmt.Println("买10个包子")
	}
}
