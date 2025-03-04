package main

import (
	"fmt"
	"strconv"
)

// 150 逆波兰表达式求值
// 给你一个字符串数组 tokens ，表示一个根据 逆波兰表示法 表示的算术表达式。
//
// 请你计算该表达式。返回一个表示表达式值的整数。
//
// 注意：
//
// 有效的算符为 '+'、'-'、'*' 和 '/' 。
// 每个操作数（运算对象）都可以是一个整数或者另一个表达式。
// 两个整数之间的除法总是 向零截断 。
// 表达式中不含除零运算。
// 输入是一个根据逆波兰表示法表示的算术表达式。
// 答案及所有中间计算结果可以用 32 位 整数表示。
func main() {
	tokens := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(tokens))
	tokens = []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(tokens))
	tokens = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(tokens))

}

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			// 取出元素，然后进行计算(弹出栈)
			// 注意取出元素的顺序
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 根据计算符计算结果
			result := 0
			switch token {
			case "+":
				result = left + right
			case "-":
				result = left - right
			case "*":
				result = left * right
			case "/":
				result = left / right
			}
			// 把结果入栈
			stack = append(stack, result)
		default:
			// 将字符串转换为整数并压入栈
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}
