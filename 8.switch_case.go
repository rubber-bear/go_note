package main

import "fmt"

func SwitchCase() {
	var age int
	fmt.Scanln(&age)
	// 注意数据类型一致
	switch age {
	case 1:
		fmt.Println("等于1")
	case 2:
		fmt.Println("等于2")
	case 3:
		fmt.Println("等于3")
	default:
		fmt.Println("都不满足")
	}
}
