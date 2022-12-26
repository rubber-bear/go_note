package main

import (
	"fmt"
	"math/big"
)

func BigInt() {
	var v1 big.Int // v1指向超大整型的空值

	v3 := new(big.Int) // 指针的方式

	// 在超大整型的值中写入值
	v1.SetInt64(1990)             // 超过64位代表的最大值会报错
	v3.SetString("324355235", 10) // 将字符串的值写入v1指向的内存地址
	fmt.Println(v3)
}
