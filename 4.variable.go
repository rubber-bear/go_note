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
变量 => 存储数据 、 存储用户输入的值

变量名要求
	变量名由字母、数字、下划线组成，且首个字符不能为数字
	不能使用Go内置的25个关键字
		- （break、default、func、interface、select、case、defer、go、map、struct、chan、
			else、goto、package、switch、const、fallthrough、if、range、type、continue、for、
			import、return、var）

全局变量和局部变量


*/

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

	var sm = "han"
	fmt.Println(sm)

}
