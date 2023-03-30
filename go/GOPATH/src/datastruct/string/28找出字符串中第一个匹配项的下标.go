package main

/*
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。
如果 needle 不是 haystack 的一部分，则返回  -1 。

示例 1：
输入：haystack = "sadbutsad", needle = "sad"
输出：0
解释："sad" 在下标 0 和 6 处匹配。
第一个匹配项的下标是 0 ，所以返回 0 。
示例 2：

输入：haystack = "leetcode", needle = "leeto"
输出：-1
解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。
*/

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 || len(needle) == 0 {
		return -1
	}

	j := 0
	n := len(needle)
	next := make([]int, n)

	for i := 0; i < len(haystack); i++ {
		for j >= 1 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}

	return -1
}

func getNext(next []int, str string) []int {
	j := 0

	for i := 1; i < len(str); i++ {
		for j > 0 && str[i] != str[j] {
			j = next[j-1]
		}
		if str[i] == str[j] {
			j++
		}
		next[i] = j
	}
	return next
}
