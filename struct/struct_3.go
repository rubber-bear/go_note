package main

import "fmt"

type Address2 struct {
	city, state string
}

type Person3 struct {
	name string
	age  int
	Address2
}

func main() {
	// 方式一：先后顺序
	var t1 = Person3{"hanYang", 30, Address2{"beijing", "china"}}
	fmt.Println(t1)

	// 方式二： 关键字
	var t2 = Person3{name: "hanL", age: 31, Address2: Address2{city: "daLian", state: "china"}}
	fmt.Println(t2.city, t2.Address2.state)

	// 方式三：先声明再赋值
	var t3 Person3
	t3.name = "zcy"
	t3.age = 26
	t3.Address2 = Address2{city: "温哥华", state: "加拿大"}
	fmt.Println(t3.city, t3.state)

}
