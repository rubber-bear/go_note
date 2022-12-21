package main

import "fmt"

// 字符串格式化

func StringFormat() {
	var name, address, action string
	// fmt.Println("字符串格式化")
	fmt.Println("请输入姓名？")
	fmt.Scanln(&name)

	fmt.Println("请输入位置？")
	fmt.Scanln(&address)

	fmt.Println("请输入行为？")
	fmt.Scanln(&action)

	result := fmt.Sprintf("我叫%s,在%s正在干%s", name, address, action)
	fmt.Println(result)
}
