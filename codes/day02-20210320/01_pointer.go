package main

import "fmt"

func main() {
	var num int = 69
	num1 := num
	//修改num1
	num1 = 99
	fmt.Println(num,num1)

	var pointerNum *int = nil
	fmt.Printf("%T,%v\n",pointerNum,pointerNum)
	//取地址
	pointerNum = &num
	fmt.Printf("%T,%v\n",pointerNum,pointerNum)
	//解引用
	*pointerNum = 116
	fmt.Println(num)    //原值被修改
	pointer2Num := &num
	fmt.Printf("%T,%v\n",pointer2Num,pointer2Num)
	*pointer2Num = 100
	fmt.Println(num)

	var pointer3Num = &num
	fmt.Printf("%T,%v\n",pointer3Num,pointer3Num)

	pp := &pointer2Num
	fmt.Printf("%T,%v\n",pp,pp)
	//new操作
	pointer := new(int)
	fmt.Printf("%T,%v,%v\n",pointer,pointer,*pointer)
	//占位符: 打印内存地址
	fmt.Printf("%p\n",pointer)
}