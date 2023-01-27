package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	协程与管道协同工作
		开启一个writeData协程，向管道中写入50个整数
		开启一个readData协程，从管道中读取writeData写入数据
		注意：writeData和readData操作的是同一个管道
		主线程需要等待writeData和readData协程否完成工作才能退出


              ------>写协程 ---
	主线程 ---｜                ｜ - > channel
              ------>写协程 ----

*/

var wg5 sync.WaitGroup

func writeData(intChan chan int) {
	defer wg5.Done()
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("写入的数据为：", i)
		time.Sleep(time.Second)
	}
	// 管道关闭
	close(intChan)
}

func readData(intChan chan int) {
	// 遍历
	defer wg5.Done()
	for v := range intChan {
		fmt.Println("读取的数据为：", v)
		time.Sleep(time.Second)
	}
}

func main() {
	// 写协程和读协程共同操作同一个管道 -> 定义管道
	intChan4 := make(chan int, 50)

	// 开启读和写的协程
	wg5.Add(2)
	go writeData(intChan4)
	go readData(intChan4)
	wg5.Wait()
}
