package main

func IsPerfectSquare(num int) bool {
	if num < 0 {
		return false
	}

	left := 0
	right := num - 1

	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// func main() {
// 	result := IsPerfectSquare(46)
// 	fmt.Println(result)
// }
