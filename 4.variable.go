package main

import "fmt"

/*
变量
	var 变量名 类型 = 值
    简化方式
 		var 变量名 = 值
    	变量名 := 值 （推荐）
	注意： 简化方式的使用前提是必须直接赋值操作
          如果是先声明再赋值，就不能使用简化方式
		// 声明变量 （如果是只声明不赋值，Go内部其实会给变量设置默认值
	    // int 为 0，float 为 0.0，bool 为 false，string 为空字符串）
        var name string
		// 变量赋值
		name= "hy"

		var (
			name = "sz"
			age = 19
			salary = 400
			hobby = "jpm"
		)


变量 => 存储数据 、 存储用户输入的值

变量名要求
	变量名由字母、数字、下划线组成，且首个字符不能为数字
	不能使用Go内置的25个关键字
		- （break、default、func、interface、select、case、defer、go、map、struct、chan、
			else、goto、package、switch、const、fallthrough、if、range、type、continue、for、
			import、return、var）


变量的作用域
	大括号中定义的变量，不能被上级调用  =》 变量的作用域就是大括号内
	大括号内可以使用上一级的变量，当前作用域（大括号），变量会先在当前作用域查找，如果没有会向上找

	变量所在的第一集大阔号内

全局变量和局部变量
	定义在全局的变量 （未写在函数中的变量） =》 全局变量
	局部有效的变量（编写在 {} 中的变量） -》 局部变量


*/
// 全局变量不能简写
// movie := "海贼王"   // 不可以
var movie string = "海贼王"

//var movie = "海贼王"

func Variable() {
	//var sb string = "hy"
	//fmt.Println(sb)
	//
	//var age int = 75
	//fmt.Println(age)
	//
	//var flag bool = true
	//fmt.Println(flag)

	//var sd string
	//sd = "hany"
	//fmt.Println(sd)

	//st := "hy"
	//fmt.Println(st)

	//var name string
	//fmt.Scanf("%s", &name)
	//fmt.Println(name)
	//
	//var sm = "han"
	//fmt.Println(sm)

	//var name, title, msg string
	//name = "st"
	//title = "st2"
	//msg = "st3"
	//fmt.Println(name, title, msg)

	//var (
	//	name    = "sz"
	//	age     = 19
	//	salary  = 400
	//	hobby   = "jpm"
	//	balance int    // 只声明但是不赋值，默认值 0
	//	gender  string // 只声明但是不赋值，又一个默认值 ""
	//	play    bool   // 只声明但是不赋值，又一个默认值 false
	//)
	//
	//fmt.Println(name, age, salary, hobby, balance, gender, play)

	//name := "wang"
	//fmt.Println(name)
	//if true {
	//	age := 18
	//	fmt.Println(age)
	//}
	//fmt.Println(name)
	// fmt.Println(age)
	fmt.Println(movie)

}
