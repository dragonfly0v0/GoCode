/*
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。

示例：

输入：s = 7, nums = [2,3,1,2,4,3] 输出：2 解释：子数组 [4,3] 是该条件下的长度最小的子数组。

提示：

1 <= target <= 10^9
1 <= nums.length <= 10^5
1 <= nums[i] <= 10^5
*/
package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	if nums == nil {
		return 0
	}

	i := 0
	sum := 0
	res := len(nums) + 1
	for j := 0; j < len(nums); j++ {
		sum += nums[j]

		fmt.Printf("sum is %v\n", sum)
		for sum >= target {
			sublength := j - i + 1
			if sublength < res {
				res = sublength
			}
			sum -= nums[i]
			i++
		}
	}
	if res == len(nums)+1 {
		return 0
	} else {
		return res
	}

}

// func main() {
// 	nums := []int{2, 3, 1, 2, 4, 3}
// 	res := minSubArrayLen(7, nums)
// 	fmt.Println(res)
// }
