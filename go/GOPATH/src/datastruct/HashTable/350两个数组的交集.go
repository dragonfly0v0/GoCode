package main

import "fmt"

/*
给你两个整数数组 nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，
应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。

示例 1：
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]

示例 2:
输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]
*/
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	dict_one := make(map[int]int)
	dict_two := make(map[int]int)
	res := []int{}

	for _, num := range nums1 {
		dict_one[num]++
	}

	for _, num := range nums2 {
		dict_two[num]++

		fmt.Printf("dict_one is %v, dict_two is %v\n", dict_one, dict_two)

		if dict_one[num] != 0 && dict_two[num] != 0 {
			res = append(res, num)
			dict_one[num]--
			dict_two[num]--
			fmt.Printf("res=%v\n", res)
		}
	}

	return res
}

func main() {
	nums1 := []int{4, 9, 5, 4}
	nums2 := []int{9, 4, 9, 8, 4}
	res := intersect(nums1, nums2)
	fmt.Println(res)
}
