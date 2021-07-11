package main

import "fmt"

func main() {
	total, incr := 0, 0
START:
	total += incr
	incr++

	if incr <= 100 {
		goto START
	}
	goto END
	fmt.Println(total) //此处不执行
END:
	fmt.Println("total:", total)
}
