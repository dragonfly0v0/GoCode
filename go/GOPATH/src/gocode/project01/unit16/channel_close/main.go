package main

import "fmt"

func main() {
	intChan := make(chan int, 10)

	intChan <- 100
	intChan <- 200
	close(intChan)

	fmt.Println("OK~")

	// 管道关闭后，读取仍然可用
	num := <-intChan
	fmt.Println(num)

	// 遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * i
	}

	close(intChan2)

	for val := range intChan2 {
		fmt.Println("v=", val)
	}
}
