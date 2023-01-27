package main

import (
	"fmt"
	"net"
)

/*
func Listen(net, addr string) (listener, error)
返回一个本地网络地址addr上监听的listener，网络类型参数net必须是面向流的网络

运行时候先启动服务端
*/

func process(conn net.Conn) {
	// 连接用完一定要关闭
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("关闭失败")
		}
	}(conn)
	for {
		// 创建一个切片，准备：将读取的数据放入切片
		buf := make([]byte, 1024)
		// 从conn连接中读取数据
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		// 将读取的内容输出
		fmt.Println(string(buf[0:n]))
	}
}

func main() {
	// 打印
	fmt.Println("服务户端启动。。。。")
	// 进行监听
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("监听失败，err", err)
		return
	}
	// 监听成功以后
	for {
		// 循环等待客户端的连接
		conn, err2 := listen.Accept()
		if err2 != nil {
			// 客户端等待失败
			fmt.Println("客户端等待失败， err2", err2)
		} else {
			fmt.Printf("等待连接成功，con=%v\n,接收到的客户端信息 %v\n", conn, conn.RemoteAddr)
		}
		// 准备一个协程，协程处理客户端服务请求
		go process(conn) // 不同的客户端请求，连接conn不一样
	}
}
