package main

import (
	"fmt"
	"unicode"
)

// 244 基本计算器
// 给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
//
// 注意:不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。
func main() {
	s := "1 + 1"
	fmt.Println(calculate(s))

	s = " 2-1 + 2 "
	fmt.Println(calculate(s))

	s = "(1+(4+5+2)-3)+(6+8)"
	fmt.Println(calculate(s))
}

func calculate(s string) int {

	stack := []int{} // 用于存储计算过程中的结果
	result := 0      // 当前的累积结果
	sign := 1        // 当前符号：1 表示正号，-1 表示负号
	n := len(s)

	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			i++
			sign = 1
		case '-':
			sign = -1
			i++
		case '(':
			// 处理左扩号
			stack = append(stack, result)
			stack = append(stack, sign)
			result = 0
			sign = 1
			i++
		case ')':
			// 乘以之前的符号
			result *= stack[len(stack)-1]
			// 加上之前的结果
			result += stack[len(stack)-2]
			// 弹出符号和结果
			stack = stack[:len(stack)-2]
			i++
		default:
			num := 0
			//for i < n && '0' <= s[i] && s[i] <= '9' {
			for i < n && unicode.IsDigit(rune(s[i])) {
				num = num*10 + int(s[i]-'0')
				i++
			}
			result += num * sign

		}
	}

	return result
}
