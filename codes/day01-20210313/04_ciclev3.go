package main

import "fmt"

func main() {
	var r2 float64
	var confirm string
	const pi2 float64 = 3.1415926

	for {
		fmt.Print("请输入半径：")
		fmt.Scan(&r2)
		fmt.Println("圆的面积是：", pi2*r2*r2)
		fmt.Print("是否继续（y/n）")
		fmt.Scan(&confirm)
		if confirm != "y" {
			break
		}
	}
}
