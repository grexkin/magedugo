package main

import "fmt"

func main() {
	score := 0
	fmt.Print("请输入成绩：")
	fmt.Scan(&score)
	switch {
	case score >= 90 && score <= 150:
		fmt.Println("A")
	case score >= 80 && score < 90:
		fmt.Println("B")
	case score >= 70 && score < 80:
		fmt.Println("C")
	case score >= 60 && score < 70:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}
}
