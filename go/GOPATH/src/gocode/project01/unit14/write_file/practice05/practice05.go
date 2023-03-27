package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//将"F:/code/test.txt"内容导入到"F:/code/jin.txt"

	// 1.读取"F:/code/test.txt"内容到内存
	// 2.将内容写入"F:/code/jin.txt"
	src_file := "F:/code/test.txt"

	//传统方法
	// src_info, err := os.Open(src_file)

	// if err != nil {
	// 	fmt.Println("Opne file err: ", err)
	// }

	//关闭文件
	// defer src_info.Close() //要及时关闭file句柄，否则会造成内存泄漏
	// reader := bufio.NewReader(src_info)

	dst_file := "F:/code/jin.txt"
	// info, err := os.OpenFile(dst_file, os.O_WRONLY|os.O_APPEND, 0666)

	// if err != nil {
	// 	fmt.Println("Opne file err: ", err)
	// 	return
	// }
	// // 及时关闭file句柄
	// defer info.Close()

	// writer := bufio.NewWriter(info)
	// for {
	// 	info, err := reader.ReadString('\n')
	// 	if err == io.EOF {
	// 		writer.WriteString(info)
	// 		break
	// 	} else if err != nil {
	// 		fmt.Println("读取错误, err: ", err)
	// 	}

	// 	writer.WriteString(info)
	// 	writer.Flush()
	// }

	src_content, err := ioutil.ReadFile(src_file)
	if err != nil {
		fmt.Println("Opne file err: ", err)
		return
	}
	err = os.WriteFile(dst_file, src_content, 0666)
	if err != nil {
		fmt.Println("Write file err: ", err)
		return
	}

}
