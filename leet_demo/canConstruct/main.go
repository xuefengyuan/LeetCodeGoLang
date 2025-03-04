package main

import "fmt"

// 383 赎金信
// 给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
//
// 如果可以，返回 true ；否则返回 false。
//
// magazine 中的每个字符只能在 ransomNote 中使用一次
func main() {
	ransomNote, magazine := "a", "b"
	ret := canConstruct(ransomNote, magazine)
	fmt.Println(ret)

	ransomNote, magazine = "aa", "ab"
	ret = canConstruct(ransomNote, magazine)
	fmt.Println(ret)

	ransomNote, magazine = "aa", "aab"
	ret = canConstruct(ransomNote, magazine)
	fmt.Println(ret)
}

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	// 创建一个字符计数器
	count := make(map[rune]int)

	// 统计 magazine 中每个字符的频率
	for _, char := range magazine {
		count[char]++
	}
	// 检查 ransomNote 中的每个字符是否能在 magazine 中找到
	for _, char := range ransomNote {
		if count[char] == 0 {
			// 如果某个字符在 magazine 中没有剩余，返回 false
			return false
		}
		// 减少该字符的频率
		count[char]--
	}

	return true
}
