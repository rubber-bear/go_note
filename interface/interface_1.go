package main

import "fmt"

type Base interface {
}

func main() {
	/*
		Go 语言中借口是一种特殊的数据类型
			type 接口名称 interface {
				方法名称() 返回值
			}
	*/

	// 定义一个切片， 内部可以存放任意类型
	// 推荐简写为 dataList := make([] interface{}, 0)
	dataList := make([]Base, 0)

	// 切片中添加字符串类型
	dataList = append(dataList, "hany")

	// 切片中添加整型
	dataList = append(dataList, 18)

	// 切片中添加 浮点型
	dataList = append(dataList, 99.33)

	fmt.Println(dataList)

}
