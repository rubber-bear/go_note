package main

import "fmt"

/*
	数组，定长且元素类型一致的数据集合
	必备的知识点
		数组的内存是连续的
		数组的内存地址实际上是数组的第一个元素的内存地址
		每个字符串的内部存储： len + "str"


	var numbers [3]int8  ------>     0 0 0    内存地址 0x.....

	var numbers := [3]int32{18, 15, 17} --------> 18 21 10

 													   无心            alex
														|                |
	numbers := [2]string{"无心"，"alex"}  --- >  str 指针(len(6))   str 指针(len(4))
												不存数据，而是存指针（存数据的位置）和长度
*/

func ArrayMemory() {

	nums := [3]int32{11, 22, 33}

	fmt.Printf("数组的内存地址：%p\n", &nums[0])
	fmt.Printf("数组的内存地址：%p\n", &nums[1])
	fmt.Printf("数组的内存地址：%p\n", &nums[2])

	names := [2]string{"无心", "hhh"}
	fmt.Printf("数组的内存地址：%p\n", &names[0])
	/*
		数组第1个元素的内存地址：0x14000136000
		数组第2个元素的内存地址：0x14000136010
	*/
	fmt.Printf("数组第1个元素的内存地址：%p\n", &names[0])
	fmt.Printf("数组第2个元素的内存地址：%p\n", &names[1])

}
