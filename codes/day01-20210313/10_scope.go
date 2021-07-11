package main

import "fmt"

var nversion string = "1.1.1"

func main() {
	fmt.Println("nversion1", nversion)
	var funcVersion string = "2.2.2"
	fmt.Println(nversion, funcVersion)
	var nversion string = "x.x.x.x"
	fmt.Println(nversion, funcVersion)
	fmt.Println("version2", nversion)
	{
		//A
		fmt.Println("version3", nversion)
		nversion = "yyyyy"
		fmt.Println("version4", nversion)
	}
	fmt.Println("version5", nversion)
	{
		//B
		fmt.Println("version6", nversion)
		var nversion string = "zzzzz"
		fmt.Println("version7", nversion)
	}
	fmt.Println("version8", nversion)
}
