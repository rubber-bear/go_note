package main

import "fmt"

type Person0 struct {
	name string
	age  int
}

var p1 Person0 = Person0{"hany", 17}
var p2 Person0 = Person0{"hanL", 16}

func doSomething1() Person0 {
	return p1
}

func doSomething2() *Person0 {
	return &p2
}

func StructBack() {
	// 结构体做参数和返回值
	// 结构体做参数和返回值时候，在执行时候会被重新拷贝一份，
	// 如果不想被拷贝，则可以通过指针的形式处理
	data := doSomething1()
	p1.name = "wan"
	fmt.Println(data)
	fmt.Println(p1)

	// 使用指针
	data2 := doSomething2()
	p2.name = "hhh"
	fmt.Println(data2) // &{hhh 16}
	fmt.Println(p2)    // {hhh 16}
}
