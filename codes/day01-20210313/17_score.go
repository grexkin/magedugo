package main

import (
	"fmt"
)

func main() {
	//成绩小练习
	var score int
	fmt.Print("请输入成绩")
	fmt.Scan(&score)
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}
	/*
		if score >=90 {
			fmt.Println("A")
		}else {
			if score >= 80 {
				fmt.Println("B")
			}else {
				if score >=70 {
					fmt.Println("C")
				}else {
					if score>=60 {
						fmt.Println("D")
					}else {
						fmt.Println("E")
					}
				}
			}
		}
	*/
}
