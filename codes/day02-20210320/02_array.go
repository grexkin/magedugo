package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var nums [10]int
	fmt.Printf("%T,%v\n",nums,nums)
	var names [10]string
	fmt.Printf("%T,%q\n",names,names)
	//字面量
	nums = [10]int{}
	fmt.Println(nums)
	nums = [10]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(nums)
	nums = [10]int{1,2,3}
	fmt.Println(nums)
	nums = [...]int{1,2,3,4,5,6,7,8,9,23}
	fmt.Println(nums)
	//nums = [...]int{1,2,3}  此处报错，这种写法必须跟定义的长度保持一致，缺一不可
	//计算字节
	fmt.Println(unsafe.Sizeof(nums))
	//关系运算
	nums2 := [3]int{1,2,3}
	nums3 := [3]int{1,2}
	fmt.Println(nums2==nums3)
	//获取长度
	fmt.Println(len(nums))
	//索引
	fmt.Println(nums[1])
	fmt.Println(nums[1:len(nums)])
	fmt.Println(nums[1:9])
	//遍历
	for i:=0;i<len(nums);i++ {
		fmt.Println(nums[i],i)
	}

	for idx,val := range nums {
		fmt.Println(idx,nums[idx],val)
	}
	//修改元素
	nums[0] = 100
	fmt.Println(nums)
}