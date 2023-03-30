package main

import "fmt"

func InsertSort(slice []int) {
	if slice == nil || len(slice) == 1 {
		return
	}

	// 0~0 0位置到0位置有序
	// 0~i 0到i也想有序
	for i := 1; i < len(slice); i++ { // 确保0~i有序
		for j := i - 1; j >= 0 && (slice[j] > slice[j+1]); j-- {
			slice[j], slice[j+1] = slice[j+1], slice[j]
		}
	}
	fmt.Println(slice)
}

func main() {
	slice := []int{9, 89, 1, 2, 10, 27, 30}
	InsertSort(slice)
}
