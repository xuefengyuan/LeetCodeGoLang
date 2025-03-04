package main

import (
	"fmt"
	"strings"
)

// 290 单词规律
// 给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。
//
// 这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律
func main() {
	pattern, s := "abba", "dog cat cat dog"
	ret := wordPattern(pattern, s)
	fmt.Println(ret)

	pattern, s = "abba", "dog cat cat fish"
	ret = wordPattern(pattern, s)
	fmt.Println(ret)

	pattern, s = "aaaa", "dog cat cat dog"
	ret = wordPattern(pattern, s)
	fmt.Println(ret)
}

func wordPattern(pattern string, s string) bool {
	// 将字符串按空格分割为单词数组
	words := strings.Split(s, " ")
	// 长度不匹配直接返回 false
	if len(words) != len(pattern) {
		return false
	}
	// pattern -> word 映射
	pTow := make(map[byte]string)
	// word -> pattern 映射
	wTow := make(map[string]byte)
	for i, word := range words {
		p := pattern[i]
		// 检查两个方向的映射是否一致
		if pTow[p] != "" && pTow[p] != word || wTow[word] > 0 && wTow[word] != p {
			return false
		}

		// 建立映射关系
		pTow[p] = word
		wTow[word] = p

	}

	return true
}
