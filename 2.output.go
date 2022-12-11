package main

import "fmt"

/*
TODO 输出
内置函数：（取标准错误中的值来输出）
	print
    println

fmt包（推荐）
	fmt Print
	fmt println
扩展：进程有stdin/stdout/stderr （标准输出，标准输入，标准错误）
*/

func Output() {
	//print("好吃不过嫂子")
	//print("好吃不过饺子")
	//println("好吃不过饺子")   // ln表示输出换行
	//println("好吃不过饺子")

	//fmt.Print("好吃不过饺子")
	//fmt.Print("好吃不过嫂子")

	//fmt.Println("好吃不过饺子")
	//fmt.Println("好吃不过嫂子")
	//fmt.Println("今天", "是个", "好日子")
	fmt.Printf("纵有%s起，%s不言弃", "疾风", "人生")
}
