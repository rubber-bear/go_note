package main

import (
	"fmt"
	"time"
)

/*
	多个协程工作，其中一个协程出现panic导致程序奔溃
	解决：refer+ recover捕获panic进行处理即使协程出现问题，主线程荏苒不受影响可以继续执行
*/

// 输出操作
func printNum() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
}

// 做除法操作
func div() {
	defer func() {
		// 调用recover内置函数可以捕获错误
		err := recover()
		if err != nil {
			fmt.Println("div 出错：", err)
		}
	}()
	num1 := 20
	num2 := 0
	res := num1 / num2
	fmt.Println(res)
}

func main() {
	// 启动两个协程
	go printNum()
	go div()
	time.Sleep(time.Second * 5)
}
