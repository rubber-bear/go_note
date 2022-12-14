package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
让用户输入数据，完成项目交互
	fmt.Scan
		既可以处理多个值，也可以知道成功几个以及错误的异常
        特别说明：要求输入两个值，如果没有输完，会一直等待
	fmt.Scanln
        基本功能与sacn一样但是
        要求输入两个值，如果有了回车，就会结束，不会等待
	fmt.Scanf
		支持模版的方式输入

无法解决的问题
	输入的信息有空格
*/

func Input() {
	/*
		// 示例 fmt.Scan
		var name string
		var age int
		fmt.Println("请输入用户名：")
		// 用户输入完成之后，会得到两个值：
		// count：用户输入了几个值，err：输入错误时候即使错误信息
		// count, err := fmt.Scan(&name, &age)  // wangcx dx
		_, err := fmt.Scan(&name, &age) // wangcx dx
		if err == nil {
			fmt.Println(name, age)
		} else {
			fmt.Println("输入错误")
			fmt.Println(err) // 1 expected integer
		}
	*/
	/*
		// 示例 fmt.Scanln
		var name string
		var age int
		fmt.Println("请输入用户名：")
		count, err := fmt.Scanln(&name, &age)
		if err == nil {
			fmt.Println(name, age)
		} else {
			fmt.Println("输入错误")
			fmt.Println(count, err)
		}
	*/

	//var name string
	//var age int
	//fmt.Println("请输入")
	//_, _ = fmt.Scanf("我叫%s 今年%d 岁", &name, &age)
	// fmt.Println(name, age)

	reader := bufio.NewReader(os.Stdin)
	// line,从stdin中读取一行数据（字节集合 -> 转化为字符串）
	// 默认一次能读取4096个字节(1032个汉字)
	// 1.一次性读完， isPrefix= false
	// 2.先读一部分， isPrefix = true, 再去赋去isPrefix = false
	line, _, _ := reader.ReadLine()
	data := string(line)
	fmt.Println(data)

}
