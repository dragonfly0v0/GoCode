package main

/*
给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：

0 <= i, j, k, l < n
nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0

示例 1：
输入：nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
输出：2
解释：
两个元组如下：
1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0

示例 2：
输入：nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
输出：1
*/
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	if len(nums1) == 0 || len(nums2) == 0 || len(nums3) == 0 || len(nums4) == 0 {
		return 0
	}

	com1 := make(map[int]int)
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			com1[v1+v2]++
		}
	}

	count := 0
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			count += com1[0-(v3+v4)]
		}
	}
	return count
}
