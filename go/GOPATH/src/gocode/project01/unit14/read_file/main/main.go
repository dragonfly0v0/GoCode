package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开一个文件
	/*
		概念说明: file的叫法
		1.file对象
		2.file指针
		3.file 文件句柄
	*/
	file, err := os.Open("F:/code/jin.txt")
	if err != nil {
		fmt.Println("Opne file err: ", err)
	}

	//关闭文件
	defer file.Close() //要及时关闭file句柄，否则会造成内存泄漏

	// 创建一个*Reader, 是带缓冲的
	//默认缓冲区是4096个字节
	/*const (
	defaultBufSize = 4096
	)
	*/
	reader := bufio.NewReader(file)
	//循环读取文件内容
	for {
		info, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Print(info)
			break
		} else if err != nil {
			fmt.Println("读取错误, err: ", err)
		}
		//输出内容
		fmt.Print(info)
	}
	fmt.Println("\n文件读取结束")
}
