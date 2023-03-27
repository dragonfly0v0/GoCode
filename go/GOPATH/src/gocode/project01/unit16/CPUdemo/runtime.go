package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 获取当前系统CPU数目
	num := runtime.NumCPU()

	runtime.GOMAXPROCS(num)
	fmt.Println("num=", num)
}
