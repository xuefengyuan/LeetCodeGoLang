package main

import "fmt"

// 242 有效的字母异位词
// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的 字母异位词
// 字母异位词是通过重新排列不同单词或短语的字母而形成的单词或短语，并使用所有原字母一次。
func main() {
	s, t := "anagram", "nagaram"
	ret := isAnagram(s, t)
	fmt.Println(ret)

	s, t = "rat", "car"
	ret = isAnagram(s, t)
	fmt.Println(ret)
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// 使用数组统计每个字符的频率
	count := make([]int, 26)
	for i := 0; i < len(s); i++ {
		// s 中字符计数加一
		count[s[i]-'a']++
		// t 中字符计数减一
		count[t[i]-'a']--
	}
	// 检查计数数组是否全部为 0

	for _, v := range count {
		if v != 0 {
			return false
		}
	}

	return true
}
