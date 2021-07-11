package main

import (
	"fmt"
	"strings"
)

func main() {
	var muti [3][5]int
	fmt.Printf("%T,%v\n",muti,muti)
	//赋值
	muti = [3][5]int{
		[5]int{1,2,3,4,5},
		[5]int{2,3,4,5,6},
		[5]int{3,4,5,6,7},
	}
	fmt.Println(muti)
	//元素访问
	fmt.Println(muti[0])
	fmt.Println(muti[0][1:4])
	//修改
	muti[0] = [5]int{9,9,9,9,9}
	fmt.Println(muti)
	//遍历
	for i:=0;i<len(muti);i++ {
		for j:=0;j<len(muti[i]);j++{
			fmt.Println(i,j,muti[i][j])
		}
	}
	fmt.Println(strings.Repeat("=",10))
	for idx,arr := range muti {
		for j,val := range arr {
			fmt.Println(idx,j,val)
		}
	}
}