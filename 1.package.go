package main

import "fmt"

/*
包管理总结
	1. 一个文件夹可以称为一个包
	2. 在文件夹（包）中可以创建多个文件
	3. 在同一个包下的每个文件中必须指定 包名称且相同
重点:
	main包，如果是main包， 则必须写一个main函数，此函数就是项目的入口，编译生成的是一个可执行文件
	非main包，将代码分门别类的，分别放在不同的包和文件中
*/

func Add() {
	fmt.Println("package.A function")
}
