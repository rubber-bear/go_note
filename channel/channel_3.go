package main

import "fmt"

func main() {
	// 定义管道 ｜｜ 声明一个管道
	var intChan3 chan int
	// 通过make初始化,管道可以存放3个int类型的数据
	intChan3 = make(chan int, 100)
	for i := 1; i < 100; i++ {
		intChan3 <- i
	}
	close(intChan3)
	// 管道的遍历 for-range
	// 注意：for-range 在遍历前要关闭管道  否则会报错：fatal error: all goroutines are asleep - deadlock!
	for v := range intChan3 {
		fmt.Println("value=", v)
	}
}
