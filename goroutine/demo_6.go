package main

import (
	"fmt"
	"sync"
	"time"
)

// RWMutex是一个读写锁，其经常用于读次数远远多余写次数的场景 （商品的读和改）
// 在读的时候，数据之间不产生影响， 写和读之间才会产生影响
var lock1 sync.RWMutex

var wg3 sync.WaitGroup // 只定义无需赋值

func read() {
	defer wg3.Done()
	lock1.RLock() // 如果是读数据，那么这个锁不产生影响，但是读写同时发生的时候，就会影响
	fmt.Println("开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("读取数据成功")
	lock1.RUnlock()
}

func write() {
	defer wg3.Done()
	lock1.Lock()
	fmt.Println("开始修改数据")
	time.Sleep(time.Second * 10)
	fmt.Println("修改数据成功")
	lock1.Unlock()
}

func main() {
	wg3.Add(6)
	// 启动协程  --- > 读多写少 读的时候会并发执行
	for i := 0; i < 5; i++ {
		go read()
	}
	go write()
	wg3.Wait()

}
