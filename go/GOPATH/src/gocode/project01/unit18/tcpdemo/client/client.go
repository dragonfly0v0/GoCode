package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.11:8888")
	if err != nil {
		fmt.Println("Dial() err= ", err)
		return
	}

	for {
		// 功能1：客户端可以发送单行数据，然后就退出
		reader := bufio.NewReader(os.Stdin) // os.Stdin, 代表标准输入[终端]

		// 从终端读取一行用户输入，并准备发送给服务器
		line, err1 := reader.ReadString('\n')
		if err1 != nil {
			fmt.Println("ReadString() err= ", err1)
		}

		// 如果输入exit就退出
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("客户端服务退出...")
			return
		}
		// 将line发给服务器
		n, err2 := conn.Write([]byte(line + "\n"))
		if err2 != nil {
			fmt.Println("conn.Write() err= ", err2)
		}
		fmt.Printf("客户端发送了 %d 个字节数据，并退出了\n", n)

	}
}
