package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
进制转换



*/

func BaseConversion() {
	// 整型转化为二进制、八进制、十六进制
	v1 := 99
	r1 := strconv.FormatInt(int64(v1), 16)
	fmt.Println(r1, reflect.TypeOf(r1))

	data := "1001000101"
	result, err := strconv.ParseInt(data, 2, 64)
	fmt.Println(result, err)
}
