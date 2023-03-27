package main

import (
	"fmt"
)

func main() {
	// 使用select解决从管道取数据的阻塞问题

	// 1.定义1个管道，10个数据 int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	// 2.定义1个管道, 5个string
	strChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 传统方法遍历管道, 如果不关闭管道会阻塞导致deadlock
	// 在实际开发中, 可能不好确定何时关闭管道
	// 可以使用select方式
label:
	for {
		select {
		case v := <-intChan: // 如果intChan一直没有关闭，不会一直阻塞而deadlock
			// 会自动到下一个case匹配
			fmt.Printf("从intChan读取数据%d\n", v)
		case v := <-strChan:
			fmt.Printf("从strChan读取数据%s\n", v)
		default:
			fmt.Println("读取结束....")
			break label
		}
	}
}
