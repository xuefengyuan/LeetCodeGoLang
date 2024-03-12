package _020_Valid_Parentheses

/**
 * @Author: darry
 * @Desc: 有效的扩号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。


示例 1：

输入：s = "()"
输出：true
示例 2：

输入：s = "()[]{}"
输出：true
示例 3：

输入：s = "(]"
输出：false


提示：

1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成
 * @Date: 2024/3/12 21:28
*/

func isValid(s string) bool {
    // 字符串为空，或者只有一个返回false
    if len(s) == 0 || len(s)%2 == 1 {
        return false
    }
    // 定义扩号的右边=左边数据
    pairs := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    stack := make([]rune, 0)
    for _, v := range s {

        if pairs[v] > 0 { // 判断，左边找不到就加入到栈里面去
            if len(stack) == 0 || stack[len(stack)-1] != pairs[v] {
                return false
            }
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, v)
        }
    }
    return len(stack) == 0
}
