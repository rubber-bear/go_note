package main

import "fmt"

type MyInt int

type Person9 struct {
	name string
	age  int
	blog string
}

func (p *Person9) do(a1 int, a2 int) int {
	return a1 + a2 + p.age
}

func (i *MyInt) DoSomething(a1 int, a2 int) int {
	return a1 + a2 + int(*i)
}

func StructType() {
	// 定义一个类型，然后为这个类型定义一些特殊的方法
	// 可以是指针，可以是类型
	// 区别是，如果不是指针v1每次会拷贝一份，指针就不会拷贝
	// 如果不使用对象，可以使用 _ 代替
	var v1 MyInt = 1
	// 只有MyInt 类型可以调用
	result := v1.DoSomething(1, 2)
	fmt.Println(result)

	//
	p1 := Person9{name: "hhh", age: 19, blog: "www.fs"}
	res := p1.do(1, 3)
	fmt.Println(res)
}
