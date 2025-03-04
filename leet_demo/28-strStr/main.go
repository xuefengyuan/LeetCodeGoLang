package main

import "fmt"

// 28.找出字符串中第一个匹配项的下标
//
// 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。如果 needle 不是 haystack 的一部分，则返回  -1
func main() {
	haystack, needle := "badbutsad", "sad"
	index := strStr(haystack, needle)
	fmt.Println(index)

	haystack, needle = "leetcode", "leeto"
	index = strStr(haystack, needle)
	fmt.Println(index)

	haystack, needle = "mississippi", "issipi"
	index = strStr(haystack, needle)
	fmt.Println(index)

	haystack, needle = "abc", "c"
	index = strStr(haystack, needle)
	fmt.Println(index)

	haystack, needle = "a", "a"
	index = strStr(haystack, needle)
	fmt.Println(index)

	haystack, needle = "ca", "a"
	index = strStr(haystack, needle)
	fmt.Println(index)

}

func strStr(haystack string, needle string) int {
	// 第1个是字符串长度，第2个是匹配的字符串长度
	n, m := len(haystack), len(needle)
	if n < m {
		return -1
	}
	// 这里把两个都是一个字母的判断考虑进去
	if n == 1 && n == m {
		return 0
	}
	// 遍历第一个字符器
	for i := 0; i < n; i++ {

		// 匹配标识，默认为true
		mark := true
		// 匹配的长度，不能超过最大字符串长度
		if i+m > n {
			break
		}
		// 遍历匹配字符串的长度
		for j := 0; j < m; j++ {
			// 匹配字符串与被匹配字符串一个一个比较，匹配字符串外层循环下标加内层循环下标，被匹配字符串是内层循环下标
			if haystack[i+j] != needle[j] {
				mark = false
				break
			}
		}
		// 匹配成功了，返回下标
		if mark {
			return i
		}
	}
	return -1
}

func strStr1(haystack string, needle string) int {
	// 获取字符串长度
	n, m := len(haystack), len(needle)
	if n < m { // 小于匹配长度，直接返回
		return -1
	}

	if n == 1 && n == m { // 这里把两个都是一个字母的判断考虑进去
		return 0
	}

	for i := 0; i < n; i++ { // 遍历第一个字符器
		mark := true // 匹配标识，默认为true
		if i+m > n { // 匹配的长度，不能超过最大字符串长度
			break
		}
		for j := 0; j < m; j++ {
			if haystack[i+j] != needle[j] { // 匹配字符串与被匹配字符串一个一个比较
				mark = false
				break
			}
		}
		if mark {
			return i
		}
	}

	return -1
}
