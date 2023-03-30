/*
给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，如果两者相等，返回 true 。# 代表退格字符。
注意：如果对空文本输入退格字符，文本继续为空。

示例 1：
输入：s = "ab#c", t = "ad#c"
输出：true
解释：s 和 t 都会变成 "ac"。

示例 2：
输入：s = "ab##", t = "c#d#"
输出：true
解释：s 和 t 都会变成 ""。

示例 3：
输入：s = "a#c", t = "b"
输出：false
解释：s 会变成 "c"，但 t 仍然是 "b"。
*/

/*
思路:
	1.双指针slow, fast, 当fast遍历到'#', slow+1=fast+1
	2.将字符串转为切片
*/

package main

func backspaceCompare(s string, t string) bool {
	// 定义双指针
	i, j := len(s)-1, len(t)-1
	// 定义函数，用于计算跳过退格符后的字符位置
	nextValidPos := func(str string, idx int) int {
		skip := 0
		for idx >= 0 {
			if str[idx] == '#' {
				skip++
				idx--
			} else if skip > 0 {
				skip--
				idx--
			} else {
				break
			}
		}
		return idx
	}
	// 从后往前遍历两个字符串
	for i >= 0 || j >= 0 {
		i = nextValidPos(s, i)
		j = nextValidPos(t, j)
		// 如果还没有遍历完一个或两个字符串
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
			i--
			j--
		} else if i >= 0 || j >= 0 {
			// 如果只有一个字符串遍历完了
			return false
		}
	}
	return true
}
