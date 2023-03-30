package main

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	res := []int{}
	val_dict := make(map[int]int)

	for _, num := range nums1 {
		val_dict[num] = 1
	}

	for _, num := range nums2 {
		if val_dict[num] == 1 {
			res = append(res, num)
			val_dict[num]--
		}
	}

	return res
}
