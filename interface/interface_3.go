package main

import (
	"fmt"
)

type IBase interface {
	f1() int
}

type Car struct {
	name string
}

func (c Car) f1() int {
	return 123
}

type User struct {
	name string
}

func (u User) f1() int {
	return 666
}

func do(base IBase) {
	// 基于接口的参数，可以实现传入多种类型(多态)
	//也同时具有约束对象必须实现接口方法的功能
	result := base.f1() // 直接调用 接口.f1() -> 找到其对应的类型并执行其方法
	fmt.Println(result)
}

func main() {
	/*
		空接口，非空接口都可以代指任何类型

	*/

	t1 := Car{name: "bl"}
	t2 := User{name: "hany"}

	do(t1)
	do(t2)

}