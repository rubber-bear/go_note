package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func PointerDetail() {
	// 1. 指针是一种数据类型, 用来表示数据的内存地址

	var v1 *string
	fmt.Println(v1, reflect.TypeOf(v1))

	// 声明一个 字符串类型的变量值为 hany
	var name string = "hany"
	// 获取name的内存地址
	fmt.Println(&name, reflect.TypeOf(&name))

	var age int = 18
	var x *int = &age
	fmt.Println(x)

	// 2. 指针存在的意义
	// 创建了一个地址的引用，以后根据这个引用再去获取它里面的值
	p1 := "hany"
	p2 := &p1
	fmt.Println(p1, p2, *p2) // hany 0x14000110340 hany

	p1 = "hanL"
	fmt.Println(p1, p2, *p2) // hanL 0x14000110340 hanL

	// 3 - 指针的使用场景1

	t1 := "zjl"
	t2 := t1
	t1 = "hany"
	fmt.Println(t1, t2) //hany zjl

	t3 := "zcy"
	t4 := &t3
	t3 = "lj"
	fmt.Println(t3, *t4)

	// 指针的使用场景2
	// 作为函数参数、 input 输入
	// var username string
	// 提示用户输入，当用户输入以后，将舒服的值赋给内存地址对应的区域中
	// fmt.Scanf("%s", &username)

	// 4- 指针的指针
	s := "hany"
	// 声明一个指针类型的变量n1，内部存储s的内存地址
	var n1 *string = &s

	// 声明一个指针的指针类型变量n2, 内部存储指针的内存地址
	var n2 **string = &n1

	// 声明一个指针的指针类型变量n3，内部存储指针n2的内存地址
	var n3 ***string = &n2

	fmt.Println(s, &s)
	fmt.Println(n1, &n1)
	fmt.Println(n2, &n2)
	fmt.Println(n3, &n3)

	// 5 - 特殊操作
	// 数组的地址 == 数组的第一个元素的地址
	dataList := [3]int8{11, 22, 33}
	fmt.Println(&dataList, reflect.TypeOf(&dataList))
	fmt.Println(&dataList[0], reflect.TypeOf(&dataList[0]))

	// 指针计算
	// 1 - 获取数组第一个元素的地址（指针）
	var firstDataPtr *int8 = &dataList[0]

	// 2 - 转换成Pointer类型
	ptr := unsafe.Pointer(firstDataPtr)

	// 3.转换成uintptr类型，然后进行内存地址的计算（即，地址增加一个字节，意味着取第二个索引位置的值）
	targetAddress := uintptr(ptr) + 1

	// 4. 根据新地址，重新转换成Pointer类型
	newPtr := unsafe.Pointer(targetAddress)

	// 5. Pointer 对象转换成为int8指针类型
	value := (*int8)(newPtr)

	// 6. 根据指针获取值
	fmt.Println("res", *value)
}
