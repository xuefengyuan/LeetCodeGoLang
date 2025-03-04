package main

import "fmt"

// 20 有效扩号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
func main() {

	s := "()"
	fmt.Println(isValid(s))

	s = "()[]{}"
	fmt.Println(isValid(s))

	s = "(]"
	fmt.Println(isValid(s))

	s = "([])"
	fmt.Println(isValid(s))

}

func isValid(s string) bool {

	// 定义一个映射，用于匹配右括号和左括号
	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// 使用一个栈来保存左括号
	stack := []rune{}

	// 遍历字符串中的每一个字符
	for _, char := range s {
		// 如果是右括号
		if ching, ok := mapping[char]; ok {
			// 如果栈为空或者栈顶元素不匹配当前右括号
			if len(stack) == 0 || stack[len(stack)-1] != ching {
				return false
			}
			// 弹出栈顶元素
			stack = stack[:len(stack)-1]
		} else {
			// 如果是左括号，压入栈中
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
