/*
给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3 输出: [ [ 1, 2, 3 ], [ 8, 9, 4 ], [ 7, 6, 5 ] ]
*/
package main

import "fmt"

/*
	1.确定矩阵的边界x, y
	2.我们规定每次遍历都是从第1个元素到该行的倒数第2个元素
	3.每次大循环都完成一圈，因此大循环的边界是n / 2, n是元素个数

*/

func generateMatrix(n int) [][]int {
	startx, starty := 0, 0
	offset := 1
	count := 1
	loop := n / 2
	mid := n % 2
	var i, j int

	nums := make([][]int, n)

	for i := 0; i < n; i++ {
		nums[i] = make([]int, n)
	}

	for loop > 0 {
		i = startx
		j = starty

		//行数不变i=startx 列数在变
		for j = starty; j < n-offset; j++ {
			nums[startx][j] = count
			count++
		}

		//行数在变 列数不变是j=n-offset
		for i = startx; i < n-offset; i++ {
			nums[i][j] = count
			count++
		}

		//行数不变是i = n-offset 列数在变
		for ; j > starty; j-- {
			nums[i][j] = count
			count++
		}

		//行数在变 列数不变j = starty
		for ; i > startx; i-- {
			nums[i][j] = count
			count++
		}
		startx++
		starty++
		offset++
		loop--
	}

	if mid == 1 {
		nums[n/2][n/2] = count
	}
	return nums
}

func main() {
	n := 4
	fmt.Println(generateMatrix(n))
}
