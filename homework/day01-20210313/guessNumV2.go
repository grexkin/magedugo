package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	guess_num := rand.Intn(100)
	guess_count := 0
	fmt.Println(guess_num)
	for {
		var input_num int
		var ack string
		fmt.Print("请输入要猜的数字：")
		fmt.Scan(&input_num)
		if input_num > guess_num {
			fmt.Printf("你猜测的数字：%d太大了\n",input_num)
			guess_count ++
		}else if input_num < guess_num {
			fmt.Printf("你猜测的数字：%d太小了\n",input_num)
			guess_count ++
		}else {
			fmt.Printf("恭喜你猜对了，%d\n",guess_num)
			fmt.Print("是否继续！Y|N")
			fmt.Scan(&ack)
			if ack == "Y" || ack == "y" {
				continue
				guess_count = 0
			} else if ack == "N" || ack == "n" {
				break
				guess_count = 5
			} else {
				continue
			}
		}
		if guess_count == 5 {
			break
		}
	}
}