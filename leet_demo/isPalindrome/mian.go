package main

import (
	"fmt"
)

// 125 验证回文串
// 如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
//
// 字母和数字都属于字母数字字符。
//
// 给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false
func main() {
	s := "A man, a plan, a canal: Panama"
	ret := isPalindrome(s)
	fmt.Println(ret)

	s = "race a car"
	ret = isPalindrome(s)
	fmt.Println(ret)

	s = " "
	ret = isPalindrome(s)
	fmt.Println(ret)
}

func isPalindrome(s string) bool {

	// 初始化左右两个指针，左指针从0开始，右指针从字符串长度-1开始
	left, right := 0, len(s)-1
	// 开始遍历，左边小于右边就继续，左边大于右边就结束循环
	for left < right {

		// 指针从左往右移动，跳过所有非字母数字字符。注意下标，左边指针要比右边指针小
		for left < right && !isNum(s[left]) {
			left++
		}
		// 指针从右往左移动，跳过所有非字母数字字符。注意下标，左边指针要比右边指针小
		for left < right && !isNum(s[right]) {
			right--
		}
		// 比较左右两个指针下标的字符，忽略大小写，不相等就返回不是回文串
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}
		// 左指针加1，右指针减1
		left++
		right--
	}

	return true
}

// 判断是不是大小写字母，数字
func isNum(b byte) bool {
	return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')
}

// 转换为小写
// 如果是小写字母，直接返回，是大写字母转换成小写字母
func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		// 大写字母加上(小写字母a减去大写字母A之间的值)
		return b + ('a' - 'A')
	}
	return b
}
