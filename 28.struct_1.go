package main

import "fmt"

/*
	结构题是一个复合类型，用于表示一组数据
	结构体由一系列属性组成，每个属性都有自己的类型和值


	type Address struct {
		city, state string
	}

	type Person struct {
		name  string
		age   int
		email string
		add   Address // 结构体嵌套
	}

	type Person2 struct {
		name    string
		age     int
		Address    // 匿名字段 //Address //Address
	}
*/

type Person struct {
	name  string
	age   int
	hobby []string
}

func structFunc() {
	// 方式一: 先后顺序
	var p1 = Person{"hany", 30, []string{"chi", "lc"}}
	fmt.Println(p1, p1.name, p1.age, p1.hobby)

	// 方式二：关键字
	var p2 = Person{name: "hanL", age: 31, hobby: []string{"tata", "kj"}}
	fmt.Println(p2)

	// 方式三: 先声明再赋值
	var p3 Person
	p3.name = "zcy"
	p3.age = 26
	p3.hobby = []string{"hj", "cy"}
	fmt.Println(p3)

}
