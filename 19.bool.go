package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
	布尔类型, 表示真假，一般是和条件等配合使用，用于满足某个条件时，执行某个操作

*/

func Bool() {
	// 字符串 转换 布尔类型
	result, err := strconv.ParseBool("true")
	fmt.Println(result, err)

	// 布尔类型 转换 字符串
	res2 := strconv.FormatBool(false)
	fmt.Println(res2, reflect.TypeOf(res2))

}
