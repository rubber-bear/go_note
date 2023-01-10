package main

import "fmt"

type f1 func(arg int) (int, bool)

// 定义函数 num1 := 1
func add(num1 int, num2 int) (int, bool) {
	result := num1 + num2
	return result, true
}

func SendEmail(arg [2]int, arg2 []int, arg3 *[2]int) {
	arg[0] = 666
	arg2[0] = 999
	arg3[0] = 888
}

func add2(arg int) (int, bool) {
	return arg + 100, true
}

func proxy(data int, exec func(int) (int, bool)) int {
	d, f := exec(data)
	if f {
		return d
	} else {
		return 999
	}
}

func proxy2(data int, exec f1) int {
	d, f := exec(data)
	if f {
		return d
	} else {
		return 999
	}
}

func do(num ...int) int {
	fmt.Println(num)
	sum := 0
	for _, value := range num {
		sum += value
	}
	return sum
}

func FuncTest() {
	// 函数参数

	// 1. 函数- 参数 注意，传值时会拷贝一份数据（等同于赋值拷贝）
	data, flag := add(1, 8)
	fmt.Println(data, flag)

	dataList := [2]int{11, 22}
	dataList2 := []int{11, 22}
	dataList3 := &[2]int{44, 55}
	SendEmail(dataList, dataList2, dataList3)
	fmt.Println(dataList)  //[11 22]
	fmt.Println(dataList2) //[11 22]
	fmt.Println(dataList3) //[11 22]

	// 2.函数 作为参数

	res := proxy(123, add2)
	res2 := proxy2(123, add2)
	fmt.Println(res)
	fmt.Println(res2)

	// 3。变长参数  切片类型, 只能放在最后，且只能有一个 --------------------
	r1 := do(1, 2, 3, 4)
	r2 := do(0, 1, 1)
	fmt.Println(r1, r2)
}
