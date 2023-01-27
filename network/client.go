package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 打印
	fmt.Println("客户端启动")
	// 调用Dial函数，参数需要指定tcp协议，需要指定服务端的IP + PORT
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		// 连接失败
		fmt.Println("客户端连接失败: err", err)
		return
	}
	fmt.Println("连接成功，conn", conn)

	// 通过客户端发送单行数据，然后退出
	reader := bufio.NewReader(os.Stdin) // os.Stdin 代表终端的标准输入
	// 从终端读取一行用户输入的信息：接受到的数据以换行来表示（回车）
	str, errR := reader.ReadString('\n')
	if errR != nil {
		fmt.Println("终端输入失败，err_r", errR)
	}
	// 将str数据发送给服务器
	n, err2 := conn.Write([]byte(str))
	if err2 != nil {
		fmt.Println("连接失败， err", err)
		return
	}
	fmt.Printf("终端数据通过客户端发送成功，一共发送了%d字节数据\n, exit", n)

}
