package main

import "fmt"

/*
   【1】Golang语言面向对象编程说明：
		Golang也支持面向对象编程(OOP)，和传统的面向对象编程有区别，不是纯粹的面向对象语言。说Golang支持面向对象编程特性是比较准确的
		Golang没有类(class)，Go语言的结构体(struct)和其它编程语言的类(Class)有同等的地位，你可以理解Golang是基于struct来实现OOP特性的
		Golang面向对象编程非常简洁，去掉了传统OOP语言的方法重载、构造西数和析构函数、隐藏的this指针等等
		Golang仍然有面向对象编程的继承，封装和多态的特性，只是实现的方式和其它OOP语言不一样


	结构体是一个复合类型，用于表示一组数据
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
	// 变量名字大写，外界就可以访问
	name  string
	age   int
	hobby []string
}

func main() {
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
