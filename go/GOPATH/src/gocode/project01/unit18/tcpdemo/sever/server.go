package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	// 循环接收客户端发送的数据
	defer conn.Close() // 接受完即关闭

	for {
		// 创建1个新的切片
		buf := make([]byte, 1024)

		// conn.Read(buf)
		// 1. 等待客户端通过conn发送消息
		// 2. 如果客户端没有write[发送], 协程就会阻塞在这个地方
		fmt.Printf("服务器在等待客户端%s输入\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println("远程客户端已关闭 ")
			return
		}
		if err != nil {
			fmt.Println("conn.Read() err1= ", err)
			return
		}
		// 3. 显示客户端发送的内容到服务器终端
		fmt.Print(string(buf[:n])) // 这个地方一定要写:n, 非常重要！
	}
}

func main() {
	fmt.Println("服务器开始监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	// 1. tcp表示使用网络协议是tcp
	// 2. 0.0.0.0:8888表示本地监听8888端口
	if err != nil {
		fmt.Println("listen err= ", err)
		return
	}
	// fmt.Printf("listen suc=%T\n", listen) // listen suc=*net.TCPListener

	defer listen.Close() //演示关闭listen

	// 循环等待客户端连接
	for {
		fmt.Println("等待客户端链接...")
		conn, err1 := listen.Accept()
		if err1 != nil {
			fmt.Println("Accept() err1= ", err1)
		} else {
			fmt.Printf("Accept() suc con=%v, 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}

		// 准备起一个协程
		go process(conn)
	}

}
