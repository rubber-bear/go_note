package main

import (
	"fmt"
	"reflect"
)

/*
	反射可以做什么？
	反射可以在运行时候动态获取变量的各种信息，比如变量的类型，类别等信息
	如果反射是结构体，还可以获取到结构体本身的信息（包括结构体的字段，方法）
	通过反射，可以修改变量的值，可以调用关联的方法
	使用反射，需要import （"reflect"）

	反射相关的函数
	1）reflect.TypeOf(变量名)， 获取变量的类型，返回reflect.Type类型
	2）reflect.ValueOf(变量名)，获取变量的值，返回reflect.Value类型（reflect.Value是一个结构体，
	通过reflect.Value 可以获取到该变量的很多信息）
*/

// 空接口没有任何方法，所以可以理解为所有类型都实现了空接口，也可以理解为我们可以将任何一个变量赋给空接口
// 利用一个函数，函数的参数定义为空接口
func testReflect(i interface{}) {
	// 调用TypeOf函数 返回一个Type类型
	t := reflect.TypeOf(i)
	fmt.Println(t)
	fmt.Printf("t的类型 %T\n", t)
	// 调用ValueOf函数，返回reflect.Value类型数据
	v := reflect.ValueOf(i)
	fmt.Println(v)
	fmt.Printf("v的类型 %T\n", v)
	// 如果向获取reValue的数值， 要调用Int方法，返回v持有的有符号整数
	num1 := 80 + v.Int()
	fmt.Println(num1)

	// 将v转成空接口
	i2 := v.Interface()
	// 类型断言
	n := i2.(int)
	n2 := n + 30
	fmt.Println(n2)

}

func main() {
	// 对基本数据类型进行反射
	// 定义一个基本数据类型
	var num int = 100
	testReflect(num)
	//fmt.Println(reflect.TypeOf(num))
	//fmt.Println(reflect.ValueOf(num))
}
