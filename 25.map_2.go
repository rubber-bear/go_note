package main

import (
	"fmt"
)

func MapFunc() {
	// 1. 长度 --------------------------------------
	data := map[string]string{"n1": "hh", "n2": "hany"}
	val := len(data)
	fmt.Println(val)

	// 会根据参数值，计算出一个合适的容量
	// 一个map中包含很多桶，每个桶中可以存放8个键值对
	info := make(map[string]string, 10)
	info["n1"] = "ll"
	info["n2"] = "hh"
	// cap（info） 会报错
	fmt.Println(info, len(info))

	// 2. 增、删、改、查------------------------------------------
	data["n3"] = "mila"

	// 修改
	data["n1"] = "npc"

	//删除
	delete(data, "n2")

	// 查看
	fmt.Println(data["n1"])

	for key, value := range data {
		fmt.Println(key, value)
	}

	// 3.value中的嵌套 ------------------------------------------
	v1 := make(map[string]int)
	v2 := make(map[string]string)
	// v3 := make(map[string]...)
	v4 := make(map[string][2]int)
	v5 := make(map[string][]int)
	v6 := make(map[string]map[int]int)
	fmt.Println(v1, v2, v4, v5, v6)

	v6["n1"] = map[int]int{1: 999, 2: 444, 3: 555}
	fmt.Println(v6)

	v7 := make(map[string][2]map[string]string)
	v7["n1"] = [2]map[string]string{map[string]string{"1": "2"}, map[string]string{"4": "5"}}

	fmt.Println(v7)

	// 4.key中的嵌套 ------------------------------------------

	v8 := make(map[int]int)
	v9 := make(map[string]int)
	v10 := make(map[float32]int)
	v11 := make(map[bool]int)
	v12 := make(map[[2]int]int)
	// v13 := make(map[ []int] int)  // 错误，不可hash
	// v14 := make(map[map[int]int] int) // 错误，不可hash
	fmt.Println(v8, v9, v10, v11, v12)

	// v15 := make(map[ [2][]int ]int) //报错
	// v15 := make(map[ [2]map[string]string ]int) //报错

	// 5.变量赋值 ------------------------------------------
	m1 := map[string]string{"n1": "hany", "n2": "kk"}
	m2 := m1
	m1["n1"] = "hjk"

	fmt.Println(m1)
	fmt.Println(m2)
	//无论是否扩容都指向同一个地址
}
