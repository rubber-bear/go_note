package main

import "fmt"

/*
	管道（channel）特质介绍
		管道本身就是一个数据结构 - 队列
		数据 先进先出
		自身线程安全，多线程访问时后，不需要加锁，channel本身就是线程安全的
		管道是有类型的，一个string的管道只能存放string类型的数据

		var 变量名 chan 数据类型
		数据类型指的是管道的类型，里面放入数据的类型，管道是有类型的，int类型的管道只能写入整数int
		管道是引用类型，必须初始化才能写入数据，即make后才能使用
*/

func main() {
	// 定义管道 ｜｜ 声明一个管道
	var intChan chan int
	// 通过make初始化,管道可以存放3个int类型的数据
	intChan = make(chan int, 3)

	// 证明管道是引用类型
	fmt.Printf("inChan的值： %v\n", intChan) // 0x1400011e018

	// 向管道存放数据
	intChan <- 10
	num := 20
	intChan <- num
	intChan <- 40
	// 不能存放大于容量的数据

	// 从管道中读取数据
	num1 := <-intChan
	num2 := <-intChan
	num3 := <-intChan

	// 注意，在没有使用协程的情况下，如果管道的数据已经全部取出，那么再取就会报错

	// 输出管道的长度
	fmt.Printf("管道的实际长度是: %v, 管道的容量是：%v\n", len(intChan), cap(intChan))
	fmt.Println(num1, num2, num3)
}
