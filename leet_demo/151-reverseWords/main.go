package main

import (
	"fmt"
	"strings"
)

// 151.反转字符串中的单词
// 给你一个字符串 s ，请你反转字符串中单词 的顺序。
// 单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
// 返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。
// 注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格
func main() {
	s := "the sky is blue"
	ret := reverseWords(s)
	fmt.Println(ret)

	s = "  hello world  "
	ret = reverseWords(s)
	fmt.Println(ret)

	s = "a good   example"
	ret = reverseWords(s)
	fmt.Println(ret)
}

// 系统api方式实现
func reverseWords(s string) string {
	// 去掉首尾空格，并用空格分割单词
	strs := strings.Fields(s)
	// 反转单词数组
	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
		strs[i], strs[j] = strs[j], strs[i]
	}
	// 将单词数组重装拼接为一个字符串
	return strings.Join(strs, " ")
}
