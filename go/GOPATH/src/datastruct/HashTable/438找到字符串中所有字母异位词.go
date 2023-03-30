package main

import "fmt"

/*
	给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
	异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

	示例 1:
	输入: s = "cbaebabacd", p = "abc"
	输出: [0,6]
	解释:
	起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
	起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。

 	示例 2:
	输入: s = "abab", p = "ab"
	输出: [0,1,2]
	解释:
	起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
	起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
	起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
*/
func findAnagrams(s string, p string) []int {
	if s == "" || p == "" {
		return []int{}
	}

	res := []int{}
	freqP := make(map[rune]int)

	for _, ch := range p {
		k := rune(ch)
		freqP[k]++
	}
	fmt.Printf("===freqP is %v\n", freqP)

	freqS := make(map[rune]int)
	left := 0
	right := 0
	for _, ch := range s {
		i := rune(s[right])
		freqS[i]++
		right++

		fmt.Printf("freqS is %v\n", freqS)

		for freqS[rune(ch)] > freqP[rune(ch)] {
			k := rune(s[left])
			fmt.Printf("freqS[rune(ch)] > freqP[rune(ch)], freqS is %v\n", freqS)
			freqS[k]--
			fmt.Printf("freqS[%v]--, freqS is %v\n", k, freqS)
			left++
		}

		if right-left == len(p) {
			fmt.Printf("left is %v\n", left)
			res = append(res, left)
		}
	}
	return res
}

// func main() {
// 	s := "cbaebabacd"
// 	p := "abcd"

// 	res := findAnagrams(s, p)
// 	fmt.Println(res)
// }
