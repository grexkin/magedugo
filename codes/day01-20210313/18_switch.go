package main

import "fmt"

func main() {
	confirm := ""
	fmt.Println("我要10个包子")
	fmt.Scan(&confirm)

	switch confirm {
	case "y":
		fmt.Println("买一个西瓜")
	case "n":
		fmt.Println("买10个包子")
	default:
		fmt.Println("买10个包子")
	}
}
