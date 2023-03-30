package main

// 要求: 求非负整数x的平方根, 返回值类型是整数，将小数部分舍去
// 输入x=4, 输出2; 输入x=8, 输出2

// 思路: 二分查找求近似
func Sqrt(num int) int {
	if num < 0 {
		return -1
	}

	left := 0
	right := num - 1
	res := 0

	for left <= right {
		mid := left + (right-left)/2

		if float64(mid*mid) <= float64(num) {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return res
}

// func main() {
// 	result := Sqrt(15)
// 	fmt.Println(result)
// }
