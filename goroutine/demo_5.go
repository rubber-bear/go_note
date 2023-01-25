package main

import (
	"fmt"
	"sync"
)

// 解决多个协程操作同一份数据
// 解决机制：确保：一个协程在执行逻辑时候另外的协程不执行
// 锁的机制 -》 加入互斥锁
/*
	type Mutex
		func (m * Mutex) Lock()
		func (m * Mutex) UnLock
*/
var lock sync.Mutex

var totalNum2 int
var wg2 sync.WaitGroup // 只定义无需赋值
func add2() {
	defer wg2.Done()
	for i := 0; i < 10000; i++ {
		// 加锁 (互斥锁)
		lock.Lock()
		totalNum2++
		// 解锁
		lock.Unlock()
	}
}

func sub2() {
	defer wg2.Done()
	for i := 0; i < 10000; i++ {
		lock.Lock()
		totalNum2--
		lock.Unlock()
	}
}
func main() {
	wg2.Add(2)
	// 启动协程
	go add2()
	go sub2()
	wg2.Wait()
	// 理论上这个totalNum结果应该是0，无论协程怎么交替执行，最终想象的结果就是0
	// 事实上：
	fmt.Println(totalNum2)

}
