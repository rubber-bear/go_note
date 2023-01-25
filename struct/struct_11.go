package main

import "fmt"

type File struct {
	fd   int
	name string
}

func NewFile(fd int, name string) *File {
	return &File{fd, name}
}

func main() {
	// 适合结构体简单，传入的值很少

	// 强制使用工厂方法，让结构体变为私有，工厂方法便有公有，
	// 这样强制所有代码在实例化结构体都是用工厂方法

	f := NewFile(10, "hhh")
	fmt.Println(f)
}
