package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
进制转换
	go的代码中
		十进制，整数的形式存在
		其他进制，以字符串的形式存在

		2进制，8进制，16进制 调用parseInt => 10进制
 		10进制  调用formatInt 转换为 （2,8,16 进制）


*/

func BaseConversion() {
	// 整型转化为二进制、八进制、十六进制
	//v1 := 99
	//r1 := strconv.FormatInt(int64(v1), 16)
	//fmt.Println(r1, reflect.TypeOf(r1))

	// data := "1001000101"
	// 如果转换成功，则将err设置为nil，result则永远以int64的类型返回
	// 注意：通过parseInt字符串转换为10进制时，本质上与Atoi是相同的
	//result, err := strconv.ParseInt(data, 2, 64)
	//fmt.Println(result, err)

	// 将十进制14转换成16进制的字符串
	v2 := strconv.FormatInt(14, 16)
	fmt.Println(v2, reflect.TypeOf(v2))

	// 将二进制"10011" 转换成16进制的字符串
	r2, _ := strconv.ParseInt("10011", 2, 0)
	r3 := strconv.FormatInt(r2, 16)
	fmt.Println(r3, reflect.TypeOf(r3))

}
