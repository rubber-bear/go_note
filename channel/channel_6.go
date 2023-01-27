package main

import (
	"fmt"
	"time"
)

/*
 管道阻塞：当管道只写数据，但是不读取数据就会阻塞
		但是读写频率不一致也不会阻塞

 select功能：解决多个管道的选择问题，也可以叫"多路复用"， 可以从多个管道中随机公平的选择一个来执行
		-- case后面必须进行的是io操作，不能是等值，随机取选择一个io操作
		-- default防止select被阻塞住，加入default
*/

func main() {
	// 定义一个int管道
	intChan7 := make(chan int, 1)
	go func() {
		time.Sleep(time.Second * 5)
		intChan7 <- 10
	}()

	// 定义一个string管道
	stringChan := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		stringChan <- "hany"
	}()
	fmt.Println(<-intChan7) // 本身取数据就是阻塞的

	select {
	// case 后面就是io操作
	case v := <-intChan7:
		fmt.Println("intChan:", v)
	case v := <-stringChan:
		fmt.Println("stringChan", v)
	default:
		fmt.Println("防止select阻塞")
	}
}
