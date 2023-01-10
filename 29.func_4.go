package main

import "fmt"

func done() int {
	fmt.Println("风吹")
	defer fmt.Println("函数执行完毕了") // 如果在这之前执行了return，defer就不会再执行
	defer fmt.Println("hhhhhhhhhh")     //defer 倒序执行
	fmt.Println("麦浪")
	return 666
}

func FuncLate() {
	// 函数延迟执行
	ret := done()
	fmt.Println(ret)

	// 自执行函数 - > 匿名函数实现
	res := func(arg int) int {
		return 100 + arg
	}(1111)
	fmt.Println(res)
}
