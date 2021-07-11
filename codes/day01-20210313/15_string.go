package main

import "fmt"

func main() {
	var desc string = "a\nb\a"
	var desc2 string = "a\\nb"
	var raw string = `
    a\nb
    fff
dasdas
`
	fmt.Println(desc)
	fmt.Println(desc2)
	fmt.Println(raw)
	//拼接
	fmt.Println("hello " + "world!")
	fmt.Println(desc2[2])
	fmt.Printf("%T,%c\n", desc2[2], desc2[3])
	desc3 := "我的家乡在" + "日喀则"
	fmt.Println(desc3[2:3]) //非ascii码会出现乱码
}
