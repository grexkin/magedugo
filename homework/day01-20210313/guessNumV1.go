package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100)
	for {
		var guess int
		fmt.Print("请输入你要猜的数字：")
		fmt.Scan(&guess)
		if guess > num {
			fmt.Printf("你猜测的数字：%d太大了\n",guess)
		}else if guess < num {
			fmt.Printf("你猜测的数字：%d太小了\n",guess)
		}else {
			fmt.Printf("恭喜你猜对了，%d\n",num)
			break
		}
	}
}