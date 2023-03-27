package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行参数有: ", len(os.Args))

	// 遍历os.Args切片，就可以得到所有的命令行输入参数值
	for k, v := range os.Args {
		fmt.Printf("args[%v]是%v\n", k, v)
	}
}
