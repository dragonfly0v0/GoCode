/*
你正在探访一家农场，农场从左到右种植了一排果树。这些树用一个整数数组 fruits 表示，其中 fruits[i] 是第 i 棵树上的水果 种类 。

你想要尽可能多地收集水果。然而，农场的主人设定了一些严格的规矩，你必须按照要求采摘水果：

你只有 两个 篮子，并且每个篮子只能装 单一类型 的水果。每个篮子能够装的水果总量没有限制。
你可以选择任意一棵树开始采摘，你必须从 每棵 树（包括开始采摘的树）上 恰好摘一个水果 。采摘的水果应当符合篮子中的水果类型。
每采摘一次，你将会向右移动到下一棵树，并继续采摘。
一旦你走到某棵树前，但水果不符合篮子的水果类型，那么就必须停止采摘。
给你一个整数数组 fruits ，返回你可以收集的水果的 最大 数目。

示例 1：
输入：fruits = [1,2,1]
输出：3
解释：可以采摘全部 3 棵树。

示例 2：
输入：fruits = [0,1,2,2]
输出：3
解释：可以采摘 [1,2,2] 这三棵树。
如果从第一棵树开始采摘，则只能采摘 [0,1] 这两棵树。
*/

/*
思路: 滑动数组
*/
package main

import "fmt"

// func Violence(fruits []int) int {

// }

func totalFruit(fruits []int) int {
	if fruits == nil {
		return 0
	}

	start := 0
	ans := 0
	count := map[int]int{}

	for k, v := range fruits {
		count[v]++

		fmt.Printf("现在count是%v\n", count)
		for len(count) > 2 {
			y := fruits[start]
			fmt.Printf("count[%v] is %v\n", y, count[y])
			count[y]--

			if count[y] == 0 {
				fmt.Printf("弹出count[%v] : %v\n", y, count[y])
				delete(count, y)
			}
			start++
		}

		if ans < k-start+1 {
			ans = k - start + 1
		}
		fmt.Println("ans现在是: ", ans)
		fmt.Println("======")
	}

	return ans
}

// func main() {
// 	fruits := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
// 	res := totalFruit(fruits)
// 	fmt.Println(res)
// }
