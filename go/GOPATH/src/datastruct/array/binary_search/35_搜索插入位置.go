package main

import "fmt"

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
你可以假设数组中无重复元素。

示例 1:
输入: [1,3,5,6], 5
输出: 2

输入: [1,3,5,6], 7
输出: 4
*/

func Search(slice []int, target int) int {
	fmt.Println("slice is ", slice)
	if slice == nil {
		return 0
	}

	left := 0
	right := len(slice) - 1
	mid := 0

	// 暴力破解
	// for i := 0; i < len(slice); i++ {
	// 	if target <= slice[i] {
	// 		return i
	// 	} else {
	// 		return len(slice)
	// 	}
	// }
	for left <= right {
		mid = int(left + (right-left)/2)
		if target == slice[mid] {
			return mid
		} else if target < slice[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right + 1
}

// func main() {
// 	slice := []int{-2, 2, 10, 30, 230}
// 	res := Search(slice, 300)
// 	fmt.Println(res)
// }
