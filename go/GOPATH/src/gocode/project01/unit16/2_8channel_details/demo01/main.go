package main

import "fmt"

func main() {
	// 管道可以声明为只读或者只写

	// 1. 默认情况，管道是双向的
	// var chann1 chan int 可读可写

	// 2.声明为只写
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	chan2 <- 20

	// 3.声明为只读
	var chan3 <-chan int
	chan3 = make(chan int, 3)
	num2 := <-chan3
	fmt.Println(num2)
}
