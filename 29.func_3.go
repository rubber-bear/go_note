package main

import "fmt"

func FuncPackage() {
	// eg
	var functionList []func()
	var functionList2 []func()

	for i := 0; i < 5; i++ {
		function := func() {
			fmt.Println(i)
		}
		functionList = append(functionList, function)
	}

	functionList[0]() // 5
	functionList[1]() // 5
	functionList[2]() // 5

	// 闭包 (通过一个函数将数据存储到函数里，以便于后续使用)
	for i := 0; i < 5; i++ {
		function := func(arg int) func() {
			return func() {
				fmt.Println(arg)
			}
		}(i)
		functionList2 = append(functionList2, function)
	}
	functionList2[0]() // 0
	functionList2[1]() // 1
	functionList2[2]() // 2

}
