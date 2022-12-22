package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
有符号整数
	int8(-128 -> 127)
	int16 (-32768 - > 32767)
	int
		在32位的操作系统上使用32位 （- 2**31 - 2**31-1 ）
        在64位的操作系统丧使用64位 （- 2** 63 - 2** 63 -1）
无符号整数
	uint8 (0-255)
	uint16 (65535)
	unit
		在32位的操作系统上使用32位 （0 - 2**32 -1）
        在64位的操作系统丧使用64位 （0- 2**64 -1）

*/

func Integer() {
	//var a uint = 32
	//fmt.Println(a)

	// 整型之间的类型转换
	//var v1 int8 = 10
	//var v2 int16 = 18
	//
	//v3 := int16(v1) + v2
	//fmt.Println(v3, reflect.TypeOf(v3))

	// 注意： 低位转向高位， 没问题
	//       高位转向低位， 需要注意
	//var v3 int16 = 130
	//v4 := int8(v3)
	//fmt.Println(v4) // -127

	// 整型与字符串之间的转换
	/*
		v5 := 19
		result := strconv.Itoa(v5)
		fmt.Println(result, reflect.TypeOf(result))

		var v6 int8  = 17
		res2 := strconv.Itoa(int(v6))
		fmt.Println(res2)
	*/

	// 字符串转换为整型
	v7 := "SB"
	res3, err := strconv.Atoi(v7)
	fmt.Println(res3, err, reflect.TypeOf(res3))

}
