package main

import (
	"fmt"
	"reflect"
)

/*
浮点型：
	go 语言中提供了两种浮点型
	float32： 用32位（4个字节）来存储浮点型
	float64： 用64位（8个字节）来存储浮点型


	float 类型，计算集中的小数的非精确的表示方式
	浮点型的底层存储原理：

	var price float32 = 39.29
		1- 第一步，浮点型转换为二进制
            整数部分，直接转换位二进制（10进制转换为二进制）即 10011
			小数部分，让小数部分乘以2，结果小于1则将结果继续乘以2，结果大于1则将结果-1继续乘以2， 结果等于1 则结束
			0.29 的二进制是将每次预算的整数部分拼接  010010100011。。。。。。
            最终：39.29  10011.0100101000111......
		2 - 第二步：
			科学计数法表示
			10011.0100101000111   =》 1.00110100101000111 * 2 **5

		3 - 第三步：存储
			以float32为例：
			-------------------------------------------------------------------
            ｜                        32 位										｜
			｜1 位 sign ｜ 8位 exponent ｜           23位	fraction					｜
			---------------------------------------------------------------------

			sign: 用1位来表示浮点数正负，0 表示正数， 1表示负数
			exponent: 指数,用8位来表示，共有256种可能（0 - 255 or -127 - 128）
					eg: 5想要存储到exponent位的话，需要让5+127 = 132，再将132转换成二进制存储到exponent（132的二进制是0100010）
			fraction:存小数点后的所有的数据


	float64和float32类似只是用于表示个部分的位数不同而已，其中 sign =1位， exponent = 11位，fraction=52位


*/

func FloatTest() {
	var v1 float32
	v1 = 3.14
	v2 := 89.33

	v3 := float64(v1) + v2
	fmt.Println(v3, reflect.TypeOf(v3))
}
