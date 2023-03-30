package main

import (
	"fmt"
)

// 1.声明函数fib
// 2.用for循环存放数列
func fib(n int) []uint64 {
	fbnSlice := make([]uint64, n)
	//第一、二个数是1
	fbnSlice[0] = 1
	fbnSlice[1] = 1

	for i := 2; i < len(fbnSlice); i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	fmt.Println(fbnSlice)
	return fbnSlice
}

func main() {
	fib(10)
	fmt.Printf("type is %T\n", fib(10))
}
