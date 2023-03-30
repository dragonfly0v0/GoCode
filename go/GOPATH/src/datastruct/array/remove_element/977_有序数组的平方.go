package main

/*
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

示例 1：
输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]

示例 2：
输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]
*/

/*
	思路:
		1.先使用快指针遍历数组, 将数组的每个数平方
		2.比较快慢指针所在数的大小, 如果nums[s] > nums[f],调换
*/

func sortedSquares(nums []int) []int {
	if nums == nil {
		return nil
	}

	left, right, k := 0, len(nums)-1, len(nums)-1

	ans := make([]int, k+1)

	for left <= right {
		lm, rm := nums[left]*nums[left], nums[right]*nums[right]

		if lm > rm {
			ans[k] = lm
			left++
		} else {
			ans[k] = rm
			right--
		}
		k--
	}

	return ans
}

// func main() {
// 	nums := []int{-4, -1, 0, 3, 10}
// 	res := sortedSquares(nums)
// 	fmt.Println(res)
// }
