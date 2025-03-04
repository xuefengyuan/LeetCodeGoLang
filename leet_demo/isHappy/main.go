package main

import (
	"fmt"
)

// 202 快乐数
// 编写一个算法来判断一个数 n 是不是快乐数。
//
// 「快乐数」 定义为：
//
// 对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
// 然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
// 如果这个过程 结果为 1，那么这个数就是快乐数。
// 如果 n 是 快乐数 就返回 true ；不是，则返回 false 。
func main() {
	n := 19
	ret := isHappy(n)
	fmt.Println(ret)

	n = 2
	ret = isHappy(n)
	fmt.Println(ret)
}

// 快慢指针
func isHappy(n int) bool {
	// 初始化慢指针和快指针
	slow, fast := n, getNext(n)
	// 如果快指针为1，或者慢指针与快指针相遇退出循环
	for slow != 0 && fast != slow {
		// 慢指针走一步
		slow = getNext(slow)
		// 快指针走两步
		fast = getNext(getNext(fast))
	}
	// 如果快指针为1，则是快乐数
	return fast == 1
}

// 计算一个数字的每位数字平方和
func getNext(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}

	return sum
}
