package main

import "fmt"

func main() {
	// 演示管道的使用
	var intChan chan int
	intChan = make(chan int, 3)

	//2. 看看intChan是什么
	// intChan的值是0xc00001e180, intChan本身的值是0xc00000a028
	fmt.Printf("intChan的值是%v, intChan本身的值是%p\n", intChan, &intChan)

	// 3.向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num

	// 4.查看管道的长度和容量
	// 管道写入数据不能超过容量
	fmt.Printf("channel len is %v, cap is %v\n", len(intChan), cap(intChan))

	// 5.从channel读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println(num2)
	fmt.Printf("channel len is %v, cap is %v\n", len(intChan), cap(intChan))

	// 6.在没有协程的情况下，若管道数据全部取出，再取就会报告deadlock

}
