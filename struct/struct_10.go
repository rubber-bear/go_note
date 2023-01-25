package main

import "fmt"

type Base struct {
	name string
}

type Son struct {
	Base // 匿名的方式继承 或者嵌套 但是： 如果改成 base //Base 就无法继承Base的方法
	age  int
}

// Base结构体的方法
func (b *Base) m1() int {
	return 666
}

// Son结构体的方法
func (s *Son) m2() int {
	return 999
}

func mian() {
	t := Son{age: 10, Base: Base{name: "han"}}
	res1 := t.m1()
	res2 := t.m2()
	fmt.Println(res1, res2)
}
