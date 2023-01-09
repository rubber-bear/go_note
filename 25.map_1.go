package main

import "fmt"

// 字典(Map)

func MapInit() {
	// userInfo := map[string]string{}
	userInfo := map[string]string{"name": "han", "age": "18"}
	fmt.Println(userInfo["name"])

	// data := make(map[int]int)
	data := make(map[int]int, 10)
	data[100] = 998
	data[200] = 998
	fmt.Println(data)

	// ------------------------用于整体赋值------------------
	// 只声明, nil
	var row map[string]int
	// row["name"] = 222
	// fmt.Println(row)  // 报错
	fmt.Println(row)

	// 声明， nil
	value := new(map[string]int)
	// value["k1"] = 123   // 会报错， 也只可以整体赋值
	// value = &data
	fmt.Println(value)

	// ------------------- 注意 -----------------
	// 键不重复 & 键必须可hash（int/ bool/ string/ array）
	v1 := make(map[[2]int]float32)
	v1[[2]int{1, 1}] = 1.6
	v1[[2]int{1, 2}] = 1.8

	v2 := make(map[[2]int][3]string)
	fmt.Println(v2)

}
