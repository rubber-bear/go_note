package main

import (
	"fmt"
	"sync"
)

// 多个协程操作同一份数据

var totalNum int
var wg1 sync.WaitGroup // 只定义无需赋值
func add() {
	// 1-取totalNum的值， totalNum + 1， 将上一个步骤的结果给totalNum
	defer wg1.Done()
	for i := 0; i < 10000; i++ {
		totalNum++
	}
}

func sub() {
	// 1-取totalNum的值， totalNum - 1， 将上一个步骤的结果给totalNum
	defer wg1.Done()
	for i := 0; i < 10000; i++ {
		totalNum--
	}
}
func main() {
	wg1.Add(2)
	// 启动协程
	go add()
	go sub()
	wg.Wait()
	// 理论上这个totalNum结果应该是0，无论协程怎么交替执行，最终想象的结果就是0
	// 事实上：
	fmt.Println(totalNum)

}
