package main

import "fmt"

func main() {
	var r float64
	const pi1 float64 = 3.1415926
	fmt.Print("请输入半径：")
	fmt.Scan(&r)
	//计算面积
	fmt.Println("圆的面积为:", pi1*r*r)
}
