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

// 两个字符串大数相加
func maxNum(s1, s2 string) string {

	// 获取两个字符串的长度
	l1 := len(s1) - 1
	l2 := len(s2) - 1
	// 结果相加的进位值
	in := 0
	// 计算的结果
	var res string

	for l1 >= 0 || l2 >= 0 || in > 0 {
		// 定义两个变量，接收两个字符串从右往左的值
		var v1, v2 int
		// 判断第一个字符串是不是已经处理完了
		if l1 >= 0 {
			v1 = int(s1[l1] - '0')
			l1--
		}

		// 判断第二个字符串是不是处理完了
		if l2 >= 0 {
			v2 = int(s2[l2] - '0')
			l2--
		}

		// 计算结果，两个字符串下标的值相加再加上进位的值
		sum := v1 + v2 + in
		// 计算个位值结果(计算结果%10=个位值)
		digit := sum % 10
		// 计算是不是有进位
		in = sum / 10
		res = fmt.Sprintf("%d%s", digit, res)

	}
	return res
}

func lengthOfLastWord(s string) int {

	// 获取单词字符串总的长度
	index := len(s) - 1
	// 跳过最后面的空格
	for s[index] == ' ' {
		index--
	}

	// 单个单词的数量
	count := 0
	// 循环统计最后一个单词的字符数量，条件：下标大于等于0并且单词下标不等于空，遇到空就退出循环
	for index >= 0 && s[index] != ' ' {
		// 单词字符个数加1
		count++
		// 下标递减
		index--
	}

	return count
}
