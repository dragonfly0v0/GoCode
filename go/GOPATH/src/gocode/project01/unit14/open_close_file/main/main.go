package main

import (
	"fmt"
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

	fmt.Printf("file is \n%v", file)

	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("Close file err: ", err)
	}
}
