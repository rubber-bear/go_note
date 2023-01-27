package main

import "fmt"

/*
	管道可以声明是只读或者只写
*/

func main() {
	// 默认情况下，管道是双向的--》可读可写
	// var intChan5 chan int

	// 声明为只写
	var intChan5 chan<- int // 管道具备只写的性质
	intChan5 = make(chan int, 3)
	intChan5 <- 30
	fmt.Println("intChan5", intChan5)

	//声明为只读
	var intChan6 <-chan int

	if intChan6 != nil {
		num := <-intChan6
		fmt.Println("num", num)
	}
	//intChan6 <- 30
}
