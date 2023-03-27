package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	/*
		4.打开1个存在的文件，将原来的内容显示在终端，并且追加五句"君子不器"
		使用os.OpenFile(), bufio.NewWriter(), *Writer的方法WriteString完成
	*/
	file := "F:/code/test.txt"
	info, err := os.OpenFile(file, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Opne file err: ", err)
		return
	}
	// 及时关闭file句柄
	defer info.Close()

	// 读取文件内容，输出到终端
	reader := bufio.NewReader(info)
	for {
		str, err := reader.ReadString('\n')
		fmt.Print(str)
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

	str := "君子不器\r\n"
	writer := bufio.NewWriter(info)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// 因为writer带缓存，因此在调用WriteString方法时
	// 内容是先写入缓存，所有需要调用Flush方法，将缓存内容真正写入文件中
	// 否则文件中会丢失数据
	writer.Flush()

}
