package main

import "fmt"

/*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。

示例 1:
输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

示例 2:
输入: strs = [""]
输出: [[""]]

示例 3:
输入: strs = ["a"]
输出: [["a"]]
*/
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	groups := make(map[[26]int][]string)

	for _, val := range strs {
		record := [26]int{}
		for _, value := range val {
			record[value-rune('a')]++
		}
		groups[record] = append(groups[record], val)
		fmt.Printf("groups[%v] is %v\n", record, groups[record])
	}
	fmt.Printf("groups is %v\n", groups)

	result := make([][]string, 0, len(groups))
	for _, val := range groups {
		fmt.Printf("result: %v append %v\n", result, val)
		result = append(result, val)
		fmt.Printf("result: %v append %v\n", result, val)
	}
	return result
}

// func main() {
// 	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
// 	res := groupAnagrams(strs)
// 	fmt.Println(res)
// }
