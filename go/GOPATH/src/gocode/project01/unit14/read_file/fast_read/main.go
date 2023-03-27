package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// ioutil
	file := "F:/code/jin.txt"
	info, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("读取错误, err: ", err)
	}
	//info此时是[]byte类型
	fmt.Printf("%v", string(info))
	// 没有显式地Open文件，因此不需要显式关闭文件
}
