package main

import "fmt"

func ArrayDetail() {
	// 数组， 定长且元素类型一致的数据集合

	// 1. 方式一：先声明在赋值， 内存中已经开辟空间，内存初始化的值都是0
	var numbers [3]int
	numbers[0] = 999
	numbers[1] = 666
	numbers[2] = 333
	fmt.Println(numbers)

	// 2.方式二：声明 + 赋值
	var names = [2]string{"无心", "笑死"}
	fmt.Println(names)

	// 方式三： 声明 + 赋值 + 指定位置
	var ages = [3]int{0: 87, 1: 45, 2: 99}
	fmt.Println(ages)

	// 方式4:省略个数
	var names2 = [...]string{"无心", "wul", "he"}
	fmt.Println(names2)

	// 声明 指针类型的数组（指针类型), 不会开辟内存初始化数组中的值 m= nil
	var m *[3]int

	// 声明数组并初始化，返回的是指针类型的数组（指针类型）
	n := new([3]int)
	fmt.Println(m, n)

}
