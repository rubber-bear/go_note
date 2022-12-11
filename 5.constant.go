package main

import "fmt"

/*
常量
   定义 const 变量名 类型 = 值
       const 变量名  =  值

	因式分解定义常量必须赋值
	iota 从0开始赋值


*/

func Constant() {
	// 定义变量
	var age1 int = 18
	age1 = 20
	fmt.Println(age1)

	// 定义常量
	const name string = "wang"
	const age2 = 19
	fmt.Println(name, age2)

	const (
		v1 = 1224
		v2 = 234343
	)
	fmt.Println(v1, v2)

	const (
		v3 = iota + 1
		_
		v4
		v5
		v6
	)
	fmt.Println(v3, v4, v5, v6)
}
