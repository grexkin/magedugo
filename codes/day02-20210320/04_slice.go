package main

import (
	"fmt"
	"strings"
)

func main() {
	var numss []int
	fmt.Printf("%T,%v,%v\n",numss,numss,numss==nil)
	//赋值
	numss = []int{1,2,3,4}
	fmt.Println(numss)
	numss = []int{1,2,3,4,5,6,7}
	fmt.Println(numss)
	numss = []int{1:100,10:8989}
	fmt.Println(numss)
	//初始化
	numss = make([]int,5)   //5各元素的切片
	fmt.Println(numss)
	numss = make([]int,5,10)  //10为容量
	fmt.Println(numss)

	numarray := [5]int{1,3,5,7,8}
	fmt.Printf("%T,%v\n",numarray[1:3],numarray[1:3])
	nums := numarray[1:5]
	fmt.Println(nums)
	//长度和容量计算
	fmt.Println(strings.Repeat("***",20))
	nums = []int{}
	fmt.Println(len(nums),cap(nums))
	nums = []int{1,2,3}
	fmt.Println(len(nums),cap(nums))
	nums = make([]int,0)
	fmt.Println(len(nums),cap(nums))
	nums = make([]int,0,10)
	fmt.Println(len(nums),cap(nums))
	nums = numarray[1:3]
	fmt.Println(len(nums),cap(nums))  //cap=len(numarray-start(1)),len=len(end-start)
	//元素的访问和修改： 通过索引来操作
	fmt.Println(nums[0])
	nums[0] = 100
	fmt.Println(nums)
	//切片中的元素遍历
	for i:=0;i<len(nums);i++ {
		fmt.Println(i,nums[i])
	}

	for idx,val := range nums {
		fmt.Println(idx,val)
	}
	//添加元素： append
	nums = append(nums,1)
	fmt.Println(nums)
	nums = append(nums,1,2,3,4)
	fmt.Println(nums)
	//删除元素
	//删除索引为0
	nums = nums[1:len(nums)]
	fmt.Println(nums)
	//删除索引为len-1
	nums = nums[0:len(nums)-1]
	fmt.Println(nums)
	//删除第i个元素
	nums1 := nums[0:1]
	nums2 := nums[2:len(nums)]
	nums = append(nums1,nums2...)  //解包操作
	fmt.Println(nums)
	//容量问题
	nums = []int{}
	fmt.Println(len(nums),cap(nums))
	nums = append(nums, 1)
	fmt.Println(len(nums),cap(nums))
	nums = append(nums, 2)
	fmt.Println(len(nums),cap(nums))
	nums = append(nums, 3)
	fmt.Println(len(nums),cap(nums))
	nums = append(nums, 4)
	fmt.Println(len(nums),cap(nums))
	nums = append(nums, 5)
	fmt.Println(len(nums),cap(nums))   //容量翻倍

	numsarray := [5]int{1,3,5,7,9}
	nums = numsarray[1:3]     //len=2,cap=4
	fmt.Println(len(nums),cap(nums))
	nums = append(nums, 1000)   //切片共享底层数组
	fmt.Println(nums)
	fmt.Println(numsarray)

	//限制容量
	//[start:end:cap_end]
	nums = numsarray[1:3:3]
	fmt.Println(len(nums),cap(nums))
	nums = append(nums,1)
	fmt.Println(nums,len(nums),cap(nums))
	fmt.Println(numsarray)
	//切片再切片
	nums = make([]int,3,10)
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	fmt.Println(nums)

	numsslice := nums[1:2]
	fmt.Println(numsslice)
	fmt.Println(len(numsslice),cap(numsslice))
	numsslice =append(numsslice,100)
	fmt.Println(nums)  //1,2,100

	//copy(dest,src)
	dest := make([]int,3)
	src :=[]int{2,3,4}
	fmt.Println(src,dest)
	copy(dest,src)
	fmt.Println(src,dest)


}
