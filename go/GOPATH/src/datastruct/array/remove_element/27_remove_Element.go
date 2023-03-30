/*
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并原地修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

示例 1: 给定 nums = [3,2,2,3], val = 3, 函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。 你不需要考虑数组中超出新长度后面的元素。

示例 2: 给定 nums = [0,1,2,2,3,0,4,2], val = 2, 函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
*/
package main

import "fmt"

/*
	思路:
	双指针法
	双指针法（快慢指针法）： 通过一个快指针和慢指针在一个for循环下完成两个for循环的工作。

	定义快慢指针

	快指针：寻找新数组的元素 ，新数组就是不含有目标元素的数组
	慢指针：指向更新 新数组下标的位置
	很多同学这道题目做的很懵，就是不理解 快慢指针究竟都是什么含义，所以一定要明确含义，后面的思路就更容易理解了。
*/

func removeElement(nums []int, val int) int {
	if nums == nil {
		return -1
	}

	// for k, v := range nums {
	// 	if v == val {
	// 		nums = append(nums[:k-1], nums[k+1:]...)
	// 	}
	// }
	// fmt.Println(nums)
	// return 1
	SlowStr := 0
	for FastStr := 0; FastStr < len(nums); FastStr++ {
		if nums[FastStr] == val {
			continue
		} else {
			nums[SlowStr] = nums[FastStr]
			SlowStr++
		}
	}
	fmt.Println(nums, SlowStr)
	return SlowStr
}

// func main() {
// 	nums := [8]int{0, 1, 2, 2, 3, 0, 4, 2}
// 	removeElement(nums[:], 2)
// }
