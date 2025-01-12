package main

import "fmt"

// 392 是子序列
// 给定两个字符串s和t，如果是的子序列true则返回，否则返回。stfalse
//
// 字符串的子序列是在不打乱原字符串相对位置的情况下，通过删除部分（也可以不删除）字符而形成的新字符串。（即是 的子"ace"序列而不是 的）。"abcde""aec"
func main() {
	s := "abc"
	t := "ahbgdc"
	ret := isSubsequence(s, t)
	fmt.Println(ret)

	s = "axc"
	t = "ahbgdc"
	ret = isSubsequence(s, t)
	fmt.Println(ret)
}

// 双指针
func isSubsequence(s string, t string) bool {
	// 定义两个字符串长度的指针
	l, r := 0, 0
	// 开始遍历，条件是两个指针小于字符串长度就结束
	for l < len(s) && r < len(t) {

		// 如果字符串匹配，s的指针向右移动
		if s[l] == t[r] {
			l++
		}
		// t的指针向右移动
		r++
	}
	// i的指针达到了s的末尾，表示s是t的子序列
	return l == len(s)
}
