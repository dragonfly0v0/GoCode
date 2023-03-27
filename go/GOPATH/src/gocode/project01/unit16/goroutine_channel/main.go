package main

import (
	"fmt"
	"sync"
)

// 1.编写函数计算阶乘， 放入map
// 2.启动多个协程，将结果放入map
// 3.map应该是全局的

var (
	myMap = make(map[int]int, 10)
	// 声明一个全局互斥锁
	// sync 是包, syncchornized
	// Mutex: 互斥
	lock sync.Mutex
)

func Step(n int) map[int]int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	lock.Lock()
	myMap[n] = res
	lock.Unlock()

	return myMap
}

func main() {
	for i := 1; i <= 80; i++ {
		go Step(i)
	}

	lock.Lock()
	// 遍历map的结果
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}
