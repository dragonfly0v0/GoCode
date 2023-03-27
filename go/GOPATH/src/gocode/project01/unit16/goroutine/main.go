package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
- 在主线程中开启goroutine，改协程每隔1秒输出"hello world"
- 在主线程每隔1秒输出"hello golang"，输出10次后退出程序
- 要求主线程和goroutine同时执行
*/

// 编写函数，每隔1秒输出"hello world"
func test() {
	for i := 1; i < 10; i++ {
		fmt.Println("test(), hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

// func main() {
// 	go test() //开启一个协程
// 	for i := 1; i < 10; i++ {
// 		fmt.Println("main(), hello golang" + strconv.Itoa(i))
// 		time.Sleep(time.Second)
// 	}
// }
