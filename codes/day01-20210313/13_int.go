package main

import "fmt"

func main() {
	var (
		i1 int = 10
		i2 int = 012
		i3 int = 0x11
		i4 int = 0b11001110
	)
	fmt.Printf("i1:%T,i1:%v\n", i1, i1)
	fmt.Printf("i2:%T,i2:%v\n", i2, i2)
	fmt.Printf("i3:%T,i3:%v\n", i3, i3)
	fmt.Printf("i4:%T,i4:%v\n", i4, i4)

	var a byte = 'A'
	fmt.Printf("%T,%v\n", a, a)

	var u rune = '大'
	fmt.Printf("%T,%v\n", u, u)

	var l int64 = 1
	fmt.Printf("%T,%v\n", l, l)

	fmt.Println(i1+i2, i1-i2, i1*i2, i1/i2, i1%i2)
	i1++
	i2--
	fmt.Println(i1, i2)
	//除法
	fmt.Println(5 / 2)

	//占位符：
	fmt.Printf("%b\n", i3)
	fmt.Printf("%b\n", 0b0010|0b1010)
	fmt.Printf("%b\n", 0b0010&^0b1010)

	i3 += 3
	fmt.Println(i3)
	fmt.Println(l + int64(i3))
	//对齐
	fmt.Printf("%+d\n", i1)
	fmt.Printf("%+5d\n", i1)
	fmt.Printf("%+05d\n", i1)
	fmt.Printf("%05d\n", i1)
	fmt.Printf("%-5d\n", i1)
}
