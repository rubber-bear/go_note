package main

import "fmt"

func Loop() {

	// 死循环
	/*
		for {
				fmt.Println("时间的复利")
				time.Sleep(time.Second * 2)
			}
	*/

	/*
		var a = 2
		for a > 0 {
			fmt.Println("时间的复利")
			a--
		}
		fmt.Printf("a=%d\n", a)
	*/

	/*
		// 变量和条件
			for i := 1; i < 10; {
				fmt.Println("纵有疾风起，人生不言弃")
				i += 1
			}
	*/

	// i ++  === i+=1
	for i := 1; i < 10; i++ {
		fmt.Println("不能够颓废呀，需要对的起家人")
	}

}
