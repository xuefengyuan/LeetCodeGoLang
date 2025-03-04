package main

import (
	"fmt"
	"sort"
)

// 49 字母异位词分组
// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//
// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))

	strs = []string{""}
	fmt.Println(groupAnagrams(strs))

	strs = []string{"a"}
	fmt.Println(groupAnagrams(strs))
}

// 排序+哈希表分组
func groupAnagrams(strs []string) [][]string {
	// 创建一个哈希表，用于存储字母异位词组
	anagrams := make(map[string][]string)
	// 遍历所有单词
	for _, str := range strs {

		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })

		// 将单词转换为字符数组并排序，得到字母异位词的标识符
		sortedStr := string(s)
		//sortedStr := sortString(str)
		// 将当前单词加入到相应的异位词组中
		anagrams[sortedStr] = append(anagrams[sortedStr], str)
	}
	// 将哈希表中的所有值（即字母异位词分组）返回
	result := [][]string{}
	for _, group := range anagrams {
		result = append(result, group)
	}

	return result
}
