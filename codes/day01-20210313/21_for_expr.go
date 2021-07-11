package main

import (
	"fmt"
)

func main() {
	var sum int
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	var totle int
	var s int = 1
	for s <= 100 {
		totle += s
		s++
	}
	fmt.Println(totle)

	var txt string
	for {
		fmt.Print("请输入内容")
		fmt.Scan(&txt)
		fmt.Println("输入的内容为：", txt)
		if txt == "2" {
			break
		}
	}
}
