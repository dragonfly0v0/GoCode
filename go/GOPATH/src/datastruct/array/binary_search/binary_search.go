package main

import (
	"fmt"
)

func InsertSort(slice []int) []int {
	for i := 1; i < len(slice); i++ {
		for j := i - 1; j >= 0 && slice[j] > slice[j+1]; j-- {
			slice[j], slice[j+1] = slice[j+1], slice[j]
		}
	}
	fmt.Println("排序后的切片是", slice)
	return slice
}

func BinarySearch(slice []int, target int) {
	// 首先对切片元素排序
	slice = InsertSort(slice)
	fmt.Println(slice)

	left := 0
	right := len(slice) - 1
	var mid int

	for left <= right {
		mid = int(left + (right-left)/2)
		if slice[mid] == target {
			fmt.Printf("target %v在排序后数组的下标是%v\n", target, mid)
			break
		} else if slice[mid] < target {
			left = mid + 1
		} else if slice[mid] > target {
			right = mid - 1
		} else {
			fmt.Printf("target %v不在切片中...", target)
		}
	}

}

// func main() {
// 	//为了每次生成的随机数不同，都需要一个seed值
// 	slice := make([]int, 6, 20)
// 	rand.Seed(time.Now().Unix())
// 	for i := 0; i < len(slice); i++ {
// 		slice[i] = rand.Intn(1000) //0 <= n < 1000
// 	}

// 	target := 0
// 	fmt.Println(slice)
// 	fmt.Println("请输入target值")
// 	fmt.Scanln(&target)
// 	BinarySearch(slice, target)

// }
