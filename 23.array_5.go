package main

import "fmt"

func ArrayNest() {
	/*
		[[0,0,0],[0,0,0]] 二维数组
	*/
	var nestData [2][3]int
	fmt.Println(nestData)

	nestData[0] = [3]int{11, 22, 33}
	nestData[1][1] = 666
	fmt.Println(nestData)

	nestData2 := [2][3]int{{11, 33, 55}, {55, 33, 22}}
	fmt.Println(nestData2)
}
