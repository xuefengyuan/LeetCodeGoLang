/**
 * @Author: darry
 * @Desc: 回文数
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数
是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

例如，121 是回文，而 123 不是。

示例 1：

输入：x = 121
输出：true
示例 2：

输入：x = -121
输出：false
解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3：

输入：x = 10
输出：false
解释：从右向左读, 为 01 。因此它不是一个回文数。

 * @Date: 2024/3/8 17:55
*/
package _009_Palindrome_Number

import (
    "fmt"
    "strconv"
)

func isPalindrome(x int) bool {
    // x % 10 取最后一位
    // x /= 10 把最后一位移除

    if x < 0 || (x%10 == 0 && x != 0) {
        return false
    }
    revertedNumber := 0
    // 取数据的一半计算
    for x > revertedNumber {

        // 中间变量 8*10 加上 传入数据的最后一位
        revertedNumber = revertedNumber*10 + x%10
        x /= 10
    }

    return x == revertedNumber || x == revertedNumber/10
}

// 121 第1位跟最后一位直接比较，相同进入下一个循环，不同则返回失败
func isPalindrome1(x int) bool {

    s := strconv.Itoa(x)
    l, r := 0, len(s)-1
    for l < r {
        fmt.Printf("l = %v, r= %v \n", l, r)
        if s[l] != s[r] {
            return false
        }
        l++
        r--
    }
    fmt.Printf("l = %v, r= %v \n", l, r)
    return true
}
