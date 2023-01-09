package main

import "fmt"

func SliceTest() {
	v1 := make([]int, 2, 5)
	v2 := append(v1, 123)

	fmt.Println(v1) // [0,0]
	fmt.Println(v2) // [0, 0, 123]

	fmt.Println(len(v1), cap(v1)) // 2, 5
	fmt.Println(len(v2), cap(v2)) // 3, 5

	v1[0] = 999
	fmt.Println(v1) // [999 0]
	fmt.Println(v2) // [999 0 123]

	v3 := make([]int, 2, 2)
	v4 := append(v3, 123)

	fmt.Println(v3, len(v3), cap(v3))
	fmt.Println(v4, len(v4), cap(v4))

	v3[0] = 999
	fmt.Println(v3) // [999, 0]
	fmt.Println(v4) // [0, 0, 123]

	v5 := [][]int{[]int{11, 22, 33, 44}, []int{44, 55}}
	v5[0][2] = 999
	fmt.Println(v5)

	v6 := [][]int{[]int{11, 22, 33, 44}, []int{44, 55}}
	v7 := v6[0]
	v7[2] = 69
	fmt.Println(v6)

	v8 := append(v6[0], 99)
	fmt.Println(v8) // [11 22 69 44 99] 扩容了

	v9 := [][]int{make([]int, 2, 5), make([]int, 2, 3)}
	v10 := append(v9[0], 99)

	fmt.Println(v9)  //[ [0, 0], [0,0] ]
	fmt.Println(v10) // [0, 0, 9]
	v10[0] = 69

	fmt.Println(v9)  //[ [69, 0], [0,0] ]
	fmt.Println(v10) // [69, 0, 99]
}
