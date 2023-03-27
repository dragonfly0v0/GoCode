package main

import "fmt"

func WriteData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Printf("WriteData()写入的数据=%v\n", i)
	}
	close(intChan)
}

func ReadData(intChan chan int, exitChan chan bool) {

	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("ReadData()读到的数据=%v\n", v)
	}
	// ReadData读取完数据后，任务完成
	exitChan <- true
	close(exitChan)
}

func main() {
	intChan := make(chan int, 100)
	exitChan := make(chan bool, 1)
	go WriteData(intChan)
	go ReadData(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
