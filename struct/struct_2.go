package main

import "fmt"

type Address struct {
	city, state string
}

type Person2 struct {
	name    string
	age     int
	address Address
}

func main() {
	// 方式一：先后顺序
	var s1 = Person2{"hanYang", 30, Address{"beijing", "china"}}
	fmt.Println(s1)

	// 方式二： 关键字
	var s2 = Person2{name: "hanL", age: 31, address: Address{city: "daLian", state: "china"}}
	fmt.Println(s2.address.city, s2.address.state)

	// 方式三：先声明再赋值
	var s3 Person2
	s3.name = "zcy"
	s3.age = 26
	s3.address = Address{city: "温哥华", state: "加拿大"}
	fmt.Println(s3.address.city, s3.address.state)

}
