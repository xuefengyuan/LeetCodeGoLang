package main

import (
	"fmt"
)

// 最后一个单词的长度
// 给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。
//
// 单词 是指仅由字母组成、不包含任何空格字符的最大子字符串
func main() {
	s1 := "4157"
	s2 := "3674"
	res := maxNum(s1, s2)
	fmt.Println(res)

	s := "   fly me   to   the moon  "
	lens := lengthOfLastWord(s)
	fmt.Println(lens)

	s = "Hello World"
	lens = lengthOfLastWord(s)
	fmt.Println(lens)

	s = "luffy is still joyboy"
	lens = lengthOfLastWord(s)
	fmt.Println(lens)

	s = "Today is a nice day"
	lens = lengthOfLastWord(s)
	fmt.Println(lens)
}

func maxNum(s1, s2 string) string {

	l1 := len(s1) - 1
	l2 := len(s2) - 1
	in := 0
	var res string

	for l1 >= 0 || l2 >= 0 || in > 0 {
		var v1, v2 int
		if l1 >= 0 {
			v1 = int(s1[l1] - '0')
			l1--
		}

		if l2 >= 0 {
			v2 = int(s2[l2] - '0')
			l2--
		}

		sum := v1 + v2 + in
		digit := sum % 10
		in = sum / 10
		res = fmt.Sprintf("%d%s", digit, res)

	}
	return res
}

func lengthOfLastWord(s string) int {

	index := len(s) - 1

	for s[index] == ' ' { // 跳过后面的空格
		index--
	}

	var count int
	for index >= 0 && s[index] != ' ' { // 统计最后一个单词的字符数量，直到遇到空格退出
		count++ // 单词字符个数
		index-- // 下标递减
	}

	return count
}
