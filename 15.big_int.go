package main

import (
	"fmt"
	"math/big"
)

func BigInt() {

	/*
		var v1 big.Int // v1指向超大整型的空值

		v3 := new(big.Int) // 指针的方式 (推荐使用)

		// 在超大整型的值中写入值
		v1.SetInt64(1990)             // 超过64位代表的最大值会报错
		v3.SetString("32435523543543535467563534556", 10) // 将字符串的值写入v1指向的内存地址
		fmt.Println(v3)

		尽量new方式取初始化并返回一个指针类型的方式
		易错的点（int 类型 和 *int指针类型）

	*/

	// 基本运算
	n1 := new(big.Int)
	n1.SetInt64(89)
	n2 := new(big.Int)
	n2.SetInt64(120)
	result1 := new(big.Int)
	result1.Add(n1, n2)
	fmt.Println(result1)

	n3 := big.NewInt(85)
	n4 := big.NewInt(145)

	result2 := new(big.Int)
	result2.Add(n3, n4)
	fmt.Println(result2)

	result3 := new(big.Int)
	//
	//result3.Sub(n3, n4)   // 减
	//result3.Mul(n3, n4)   // 乘
	//result3.Div(n3, n4)   // 除（地板除， 只能得倒商）

	minder := new(big.Int)
	result3.DivMod(n3, n4, minder)

	fmt.Println(result3, minder)

	fmt.Println(result3.String())

	n5 := new(big.Int)
	n5.SetString("324242233552525444545545", 10)

	n6 := new(big.Int)
	n6.SetString("32424223355252544454554545545", 10)

	result4 := new(big.Int)
	result4.Add(n5, n6)
	fmt.Println(result4.String())

	// Add 等方法传的参数值指针类型

}
