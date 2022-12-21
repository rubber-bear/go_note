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

	/*
		// i ++  === i+=1
			for i := 1; i < 10; i++ {
				fmt.Println("不能够颓废呀，需要对的起家人")
			}
	*/

	/*
		// continue
		for i := 1; i < 3; i++ {
				for j := 1; j < 5; j++ {
					if j == 3 {
						continue
					}
					fmt.Println(i, j)
				}
			}
	*/

	/*
		// break
		for {
			fmt.Println("start")
			break
			fmt.Println("end")
		}
	*/

	/*
		data := 66
		for {
			var inputNum int
			fmt.Println("输入数字")
			fmt.Scanln(&inputNum)
			if inputNum > data {
				fmt.Println("大了")
			} else if inputNum < data {
				fmt.Println("小了")
			} else {
				fmt.Println("猜对了")
				break
			}
		}
	*/

	/*
		// 对for 循环打标签，然后通过break和continue就可以实现多层循环的跳出和终止

		f1:
		for i := 1; i < 3; i++ {
			for j := 1; j < 5; j++ {
				if j == 3 {
					continue f1
				}
				if j == 2 {
					break f1
				}
				fmt.Println(i, j)
			}
		}
	*/

	// goto

	var name string
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)
	if name == "A" {
		goto VIP
	} else if name == "S" {
		goto SVIP
	}
	fmt.Println("预约")
VIP:
	fmt.Println("等位")
SVIP:
	fmt.Println("进入")

}
