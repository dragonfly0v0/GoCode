package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 编写函数，接收两个文件路径src_path, dst_path
func CopyFile(src_path string, dst_path string) (written int64, err error) {
	src_file, err1 := os.Open(src_path)
	if err1 != nil {
		fmt.Println("Open file error: ", err)
	}
	defer src_file.Close()

	reader := bufio.NewReader(src_file)

	// 打开dst_path
	dst_file, err2 := os.OpenFile(dst_path, os.O_WRONLY|os.O_CREATE, 0666)
	if err2 != nil {
		fmt.Println("Open file error: ", err)
		return
	}
	defer dst_file.Close()

	writer := bufio.NewWriter(dst_file)
	return io.Copy(writer, reader)
}

// 定义1个结构体，用于保存统计类型的数目
type Statics struct {
	Letter_num int
	Num        int
	Space_num  int
	Other_num  int
}

func main() {
	/*
	  思路:
	  打开一个文件，创建一个reader
	  每读取一行，就去统计改行有多少个英文、数字、空格和其他字符
	  将结果保存在一个结构体中
	*/
	src_path := "F:/code/self_intro.txt"
	src_file, err := os.Open(src_path)
	if err != nil {
		fmt.Println("Open file error: ", err)
	}
	defer src_file.Close()

	reader := bufio.NewReader(src_file)

	// 定义Statics实例
	var count Statics
	for {
		str, err1 := reader.ReadString('\n')
		if err1 == io.EOF {
			break
		}

		// str2 := []rune(str)
		// 遍历str统计
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.Letter_num++
			case v == ' ' || v == '\t':
				count.Space_num++
			case v >= '0' && v <= '9':
				count.Num++
			default:
				count.Other_num++
			}
		}
	}

	fmt.Printf("字符个数为: %v, 数字个数为: %v, 空格个数为: %v, 其他字符个数为: %v", count.Letter_num,
		count.Num, count.Space_num, count.Other_num)
}
