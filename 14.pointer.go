package main

import "fmt"

/*
	声明变量
	var v1 int  指针指向内存中的值，有默认值0
	v2 := 99    指针指向内存中的值，值为99

	var v3  *int 创建一块内存指向内存地址0x0， 0x0指向的值为nil

	v4 := new(int) 在内存中会创建两块内存  v4 指向一个内存地址， 内存地址指向值（存储数据的地方）
    new的作用，用于创建内存，用于内部数据的初始化，并返回一个指针类型

	nil： nil指go语言中的空值
*/

func Pointer() {
	fmt.Println("指针")
}
