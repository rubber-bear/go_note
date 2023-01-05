package main

import "fmt"

func ArrayFunc() {
	// 1. 长度
	name1 := [2]string{"无心", "无情"}
	fmt.Println(len(name1))

	//2. 索引
	name2 := [2]string{"无心", "无脑"}
	data := name2[0]
	fmt.Println(data)

	name2[0] = "hh"
	fmt.Println(name2)

	// 3.切片
	nums := [3]int32{11, 22, 33}
	data2 := nums[0:2]
	fmt.Println(data2)

	// 4. 循环

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	// for ... range 循环
	for index, item := range nums {
		fmt.Println(index, item)
	}

	for _, item1 := range nums {
		fmt.Println(item1)
	}

}
