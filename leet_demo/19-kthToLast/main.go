package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 19
// 找到链表中倒数第n个节点的值返回
func main() {
	//head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	//n := 2
	//fmt.Println(kthToLast(head, n)) // 输出: 4

	No := "AAAC-250226"
	serialNumbers := strings.Split(No, "-")

	serialNumber, _ := strconv.ParseInt(strings.TrimSpace(serialNumbers[1][6:]), 10, 64)
	fmt.Println(serialNumber)
}

// ListNode 定义链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 暴力解法
func kthToLast1(node *ListNode, k int) int {
	values := make([]int, 0)
	for node != nil {
		values = append(values, node.Val)
		node = node.Next
	}

	return values[len(values)-k]
}

// 滑动窗口
func kthToLast(node *ListNode, n int) int {
	if node == nil || n <= 0 { // 边界检查
		return -1
	}
	fast := node
	slow := node

	// 快指针先走n步
	for i := 0; i < n; i++ {
		if fast == nil {
			return -1 // n超出链表长度
		}
		fast = fast.Next
	}

	// 快慢指针同步移动，比如n=2，快指针在n下标的前一个向右移动，慢指针就从开始移动，一直移动到n的下标位置
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// 返回节点的值
	return slow.Val
}
