package main

import (
	"fmt"
	"time"
)

/*
waitGroup的作用
waitGroup用于等待一组线程的结束，父线程调用Add方法来设定等待的线程数量，每个被等待的线程结束时候应调用Done
方法，同时主线程里可以调用wait方法阻塞值所有线程结束，解决主线程在子线程结束后自动结束

主要方法

	func(*waitGroup) Add
	Add 方法向内部技术加上delta，delta可以是负数；如果内部计数器变为0，wait方法阻塞等待所有线程都会释放，
	如果计数器小于0，方法panic。注意Add加上正数的调用应在wait之前，否则wait可能只会等待很少的线程，一般来说
	本方法应在创建新的线程或者其他等待的事件之前调用

	func（*waitGroup）Done
	Done方法减少waitGroup计数器的值，应在线程的最后执行

	func（*waitGroup）Wait
	wait方法阻塞知道waitGroup计数器减为0
*/
func main() {
	// 当前匿名函数 + 外部变量 = 闭包
	for i := 1; i <= 5; i++ {
		go func(n int) {
			// 使用匿名函数，直接调用
			fmt.Println(n)
		}(i)
	}

	time.Sleep(time.Second * 2)
}
