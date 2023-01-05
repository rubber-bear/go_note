package main

import "fmt"

/*
	可变和拷贝
	可变：数组的元素可以被修改(长度和类型不可修改)
	拷贝：重新拷贝

		变量赋值时候重新拷贝一份
*/
func ArrayCopy() {
	name1 := [2]string{"无心", "呵呵"}
	name2 := name1

	name1[1] = "han"
	fmt.Println(name1, name2)
}
