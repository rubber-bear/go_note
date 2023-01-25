package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("hello golang + " + strconv.Itoa(i))
		// 阻塞1秒
		time.Sleep(time.Second)
	}
}

func main() {
	/*
		   程序开始，主线程开始执行
		         ｜
			go test() 开始协程  -------> 执行协程中代码逻辑...
		         ｜
			主线程代码继续执行
		         ｜
			程序结束/主线程结束


		主死随从
	*/
	go test() // go, 开启一个协程
	for i := 1; i <= 6; i++ {
		fmt.Println("hello hany + " + strconv.Itoa(i))
		// 阻塞1秒
		time.Sleep(time.Second)
	}
}
