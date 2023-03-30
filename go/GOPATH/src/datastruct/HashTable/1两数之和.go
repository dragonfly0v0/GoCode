package main

func TwoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	val_map := make(map[int]int)

	for k, v := range nums {
		s := target - v
		if other, ok := val_map[s]; ok {
			return []int{other, k}
		}
		val_map[v] = k
	}
	return []int{}
}
