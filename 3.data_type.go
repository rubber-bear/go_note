package main

import "fmt"

/*
数据类型
	整型 （ + - * /（整除） % 取余）
	字符串 引号 ""
		字符串操作：相加
	布尔 False(假), True（真）



*/

func DataType() {
	// 整型
	fmt.Println(666)
	fmt.Println(6 + 8)
	fmt.Println(6 - 8)
	fmt.Println(6 * 8)
	fmt.Println(6 / 8)
	fmt.Println(6 % 8)

	// 字符串
	fmt.Println("纵有疾风起" + "人生不言弃")

	// 布尔
	fmt.Println(1 > 2)

}
