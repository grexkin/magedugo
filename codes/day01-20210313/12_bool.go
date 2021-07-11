package main

import "fmt"

func main() {
	var isBoy bool
	fmt.Printf("%T,%v\n", isBoy, isBoy)
	t1, t2 := true, true
	f1, f2 := false, false
	fmt.Println(t1 && t2, t1 && f1, f1 && t2, f1 && f2)
	fmt.Println(t1 || t2, t1 || f1, f1 || t2, f1 || f2)
	fmt.Println(!f1, !f2)
}
