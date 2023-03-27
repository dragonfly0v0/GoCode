package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		3.打开1个存在的文件，在原来的内容下追加"吾日三省吾身"
		使用os.OpenFile(), bufio.NewWriter(), *Writer的方法WriteString完成
	*/
	file := "F:/code/test.txt"
	info, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Opne file err: ", err)
		return
	}
	// 及时关闭file句柄
	defer info.Close()

	str := "吾日三省吾身\r\n"
	writer := bufio.NewWriter(info)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	// 因为writer带缓存，因此在调用WriteString方法时
	// 内容是先写入缓存，所有需要调用Flush方法，将缓存内容真正写入文件中
	// 否则文件中会丢失数据
	writer.Flush()

}
