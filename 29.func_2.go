package main

import "fmt"

type f2 func(int, int) string

func exec(num1 int, num2 int) string {
	fmt.Println("函数执行")
	return "success"
}

func getFunction() func(int, int) string {
	return exec
}

func getFunction2() f2 {
	return exec
}

func F1(n1 int, n2 int) func(int) string {
	return func(n1 int) string {
		fmt.Println("匿名函数")
		return "匿名"
	}
}

func FuncUse() {
	// 函数返回值

	// 1 - 返回一个函数
	function := getFunction()
	res := function(111, 222)
	fmt.Println(res)

	// 2 - 匿名函数
	// --------------------
	v1 := func(n1 int, n2 int) int {
		return 123
	}
	data := v1(11, 22)
	fmt.Println(data)

	// ------------------------
	value := func(n1 int, n2 int) int {
		return 123
	}(11, 55)
	fmt.Println(value)

	// -------------------------

}
