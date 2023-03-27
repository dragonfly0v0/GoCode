package main

import "fmt"

func putNum(intChan chan int) {
	for i := 1; i < 80000; i++ {
		intChan <- i
	}

	// 关闭intChan
	close(intChan)
}

func JudgeNum(num int) bool {
	if num == 0 || num == 1 {
		return false
	}

	if num == 2 {
		return true
	}

	for i := 2; i <= (num / 2); i++ {

		if num%i == 0 {
			// fmt.Printf("i是%d, num%d不是素数\n", i, num)
			return false
		}
	}
	return true
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for {
		num, ok := <-intChan
		if !ok {
			break
		}

		// 判断num是不是素数
		res := JudgeNum(num)
		if res {
			primeChan <- num
		}
	}

	fmt.Println("primeNum获取数据完毕, 退出。。。")
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)

	// 开启1个协程，向intChan放入1-80000个数
	go putNum(intChan)

	for i := 0; i < 4; i++ {
		// 开启1个协程，从intChan取出数据，判断是否是素数，如果是，就放入primeChan
		go primeNum(intChan, primeChan, exitChan)
	}

	// 从主线程进行处理
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
		// 当从exitChan取出4个结果，就可以放心关闭primeNum
	}()

	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数=%d\n", res)
	}

	fmt.Println("主线程退出...")
}
