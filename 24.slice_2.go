package main

import (
	"fmt"
	"reflect"
)

func SliceUse() {

	// 1. 长度和容量 ------------------------------------------------
	v1 := []int{11, 22, 33}
	fmt.Println(len(v1), cap(v1))

	// 2.索引 ------------------------------------------------
	v2 := []string{"a", "b", "c"}
	fmt.Println(v2[0], v2[1])

	v3 := make([]int, 2, 5)
	fmt.Println(v3[1]) // v3[2] 会报错

	v2[2] = "d"
	fmt.Println(v2)

	// 3.切片 ------------------------------------------------
	v4 := []int{11, 22, 33, 44, 55}
	v5 := v4[1:2]
	fmt.Println(v5, reflect.TypeOf(v5)) //[22] []int

	// !!!!!!!重要：只是切出来一部分，但是存储的地址和v4一样
	v6 := v4[1:]
	v7 := v4[:4]
	fmt.Println(v6, v7)

	// 4. 追加------------------------------------------------
	v8 := []int{11, 22, 33}
	v9 := append(v8, 44, 55, 66, 77, 88)

	v10 := append(v8, []int{100, 200, 300}...)
	fmt.Println(v9, v10)

	// 5.删除 （切片一般不实用删除， 操作费劲，覆盖原来切片的内容）----------------------------------------------------------------------
	v11 := []int{11, 22, 33, 44, 55, 66}
	deleteIndex := 2
	result := append(v11[:deleteIndex], v11[deleteIndex+1:]...)
	fmt.Println(result)
	/*
		// todo !!!! ⚠️
		切片获取到 {11, 22}
		又获取到 {44, 55, 66}, 将 44，55，66要添加到切片里，由于共用存储地址 会将v11覆盖
	*/
	fmt.Println(v11) //[11 22 44 55 66 66]

	// 6。插入 (效率低下)-------------------------------------------------------------------
	v12 := []int{11, 22, 33, 44, 55, 66}
	insertIndex := 3 // 在索引位置插入99
	res := make([]int, 0, len(v1)+1)
	res = append(res, v12[:insertIndex]...)
	res = append(res, 99)
	res = append(res, v12[insertIndex:]...)

	fmt.Println(res)

	// 7.循环
	v13 := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	for i := 0; i < len(v13); i++ {
		fmt.Println(i, v13[i])
	}

	for index, value := range v13 {
		fmt.Println(index, value)
	}
}
