package main

import "fmt"

func main() {
	//二维数组
	var slice []int = []int{1, 3, 5, 8, 9, 10}
	var slice_01 []int = make([]int, 10)

	//拷贝
	copy(slice_01, slice)
	fmt.Println(slice_01)
}
