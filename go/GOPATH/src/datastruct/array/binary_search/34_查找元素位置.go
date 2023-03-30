package main

func GetLeft(nums []int, target int) int {
	// var slice = []int{-1, -1}
	// if nums == nil {
	// 	return slice
	// } else {
	// 	for i := 0; i < len(nums); i++ {
	// 		if nums[i] == target {
	// 			slice = append(slice, i)
	// 			fmt.Println(slice)
	// 		}
	// 	}
	// 	slice = append(slice[1:2], slice[len(slice)-1])
	// 	return slice
	// }

	// 情况一：target 在数组范围的右边或者左边，例如数组{3, 4, 5}，target为2或者数组{3, 4, 5},target为6，此时应该返回{-1, -1}
	// 情况二：target 在数组范围中，且数组中不存在target，例如数组{3,6,7},target为5，此时应该返回{-1, -1}
	// 情况三：target 在数组范围中，且数组中存在target，例如数组{3,6,7},target为6，此时应该返回{1, 1}
	left := 0
	right := len(nums) - 1
	mid := 0
	leftBorder := -2

	for left < right {
		mid = left + (right-left)/2
		if target <= nums[mid] {
			right = mid - 1
			leftBorder = right
		} else {
			left = mid + 1
		}
	}
	return leftBorder
}

func searchRange(nums []int, target int) []int {
	if nums == nil {
		return []int{-1, -1}
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			i, j := mid, mid
			for i >= 0 && nums[i] == target {
				i--
			}
			for j < len(nums) && nums[j] == target {
				j++
			}
			return []int{i + 1, j - 1}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return []int{-1, -1}
}

// func main() {
// 	nums := []int{1, 2, 2, 4, 5, 5, 5}
// 	res := searchRange(nums, 0)
// 	fmt.Println(res)
// }
