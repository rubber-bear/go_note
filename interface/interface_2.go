package main

import "fmt"

type Teacher struct {
	age  int
	name string
}

type Role struct {
	title string
}

func something(arg interface{}) {
	// 参数是接口类型
	// fmt.Println(arg)
	// 接口转换为Teacher
	tp, success := arg.(Teacher)
	if success {
		fmt.Println(tp.name, tp.age)
	} else {
		fmt.Println("failed")
	}
}

func doTest(arg interface{}) {
	// 多个类型转换
	switch tp := arg.(type) {
	case Teacher:
		fmt.Println(tp.name)
	case Role:
		fmt.Println(tp.title)
	case string:
		fmt.Println(tp)
	default:
		fmt.Println(tp)
	}

}

func main() {
	something("hany")
	something("444")
	something("435.99")
	something(Teacher{name: "hanL", age: 19})

	doTest(Teacher{name: "hanL", age: 19})
	doTest(Role{title: "hany_my"})
	doTest("hanL")
	doTest("4446666")
	doTest("43511.99")
}
