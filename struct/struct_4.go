package main

import "fmt"

type people struct {
	name string
	age  int
}

func main() {
	// 初始化结构体 - 创建一个结构体对象
	p := people{"xll", 29}
	fmt.Println(p.name, p.age)

	// 初始化结构体指针
	/*
		             -------------
		 	p  --->  | wang | 29 |
		             -------------



					 -------------
		             | wang | 29 |
					 -------------
				内存地址：xc0000a60200
							^
							|
					 -------------
		p0 ----> 	|oxc0000a60200|   结构体指针
					 -------------
					内存地址： xc000a640



	*/
	// var p0 *people = &people{"wang", 29}
	p0 := &people{"wang", 29}
	fmt.Println(p0, p0.name, p0.age)

	var pt *people = new(people) // 结构体指针
	pt.name = "hh"
	pt.age = 30
	fmt.Println(pt, pt.name, pt.age)

}
