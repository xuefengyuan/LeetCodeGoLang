package main

import (
	"fmt"
)

// 链表反转，遍历，递归
// 1 2 3 4 5
// 5 4 3 2 1

type ListNode struct {
	val  int
	next *ListNode
}

func main() {

	node5 := &ListNode{5, nil}
	node4 := &ListNode{4, node5}
	node3 := &ListNode{3, node4}
	node2 := &ListNode{2, node3}
	node1 := &ListNode{1, node2}
	node := reverseList(node1)
	fmt.Printf("\n = = = = %+v", node)

	node = reverseList2(node1)
	fmt.Printf("\n = = = = %+v", node)

}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	return prev
}

func reverseList2(head *ListNode) *ListNode {

	if head == nil || head.next == nil {
		return head
	}
	newHead := reverseList2(head.next)
	head.next.next = head
	head.next = nil
	return newHead
}
