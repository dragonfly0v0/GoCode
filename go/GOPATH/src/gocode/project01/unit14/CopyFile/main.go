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
		fmt.Println("Open file error: ", err1)
	}
	defer src_file.Close()

	reader := bufio.NewReader(src_file)

	// 打开dst_path
	dst_file, err2 := os.OpenFile(dst_path, os.O_WRONLY|os.O_CREATE, 0666)
	if err2 != nil {
		fmt.Println("Open file error: ", err2)
		return
	}
	defer dst_file.Close()

	writer := bufio.NewWriter(dst_file)
	return io.Copy(writer, reader)
}

func main() {

}
