package main

import "fmt"

func main() {
	// 定义管道 ｜｜ 声明一个管道
	var intChan2 chan int
	// 通过make初始化,管道可以存放3个int类型的数据
	intChan2 = make(chan int, 3)
	intChan2 <- 13
	intChan2 <- 14

	// 关闭管道 不能再次写入，然是可以读数据
	close(intChan2)
	// intChan2 <- 18

	num := <-intChan2
	fmt.Println(num)

}
