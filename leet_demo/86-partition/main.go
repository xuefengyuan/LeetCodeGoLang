package main

import "fmt"

// 86 分割链表
// 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
//
// 你应当 保留 两个分区中每个节点的初始相对位置。
func main() {
	// 创建测试链表: 1 -> 4 -> 3 -> 2 -> 5 -> 2
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 4}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 2}

	fmt.Println("原链表:")
	printList(head)

	// 分割链表，x = 3
	newHead := partition(head, 3)

	fmt.Println("分割后的链表:")
	printList(newHead)
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("null")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双链表
func partition(head *ListNode, x int) *ListNode {
	// 两个虚拟头节点
	// 用于保存小于x的节点链表
	smallHead := &ListNode{}
	// 用于保存大于于x的节点链表
	largeHead := &ListNode{}

	// 定义两个指针，分别指向上面两个虚拟头节点
	small := smallHead
	large := largeHead
	// 遍历原链表，条件是head不为空
	for head != nil {

		// 判断当前节点小于x，接入smaller链表
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			// 当前节点大于等于x，接入larger链表
			large.Next = head
			large = large.Next
		}

		// 移动到下一个节点
		head = head.Next
	}
	// 防止large链表出现环，断开最后一个节点
	large.Next = nil
	// 连接两个链表
	small.Next = largeHead.Next
	// 返回新的链表头
	return smallHead.Next
}

// 双链表
func partition1(head *ListNode, x int) *ListNode {
	// 两个虚拟头节点
	// 用于保存小于x的节点链表
	smallHead := &ListNode{}
	// 用于保存大于于x的节点链表
	largeHead := &ListNode{}

	// 定义两个指针，分别指向上面两个虚拟头节点
	small := smallHead
	large := largeHead

	// 遍历原链表，条件是head不为空
	for head != nil {
		// 判断当前节点小于x，接入smaller链表
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			// 当前节点大于等于x，接入larger链表
			large.Next = head
			large = large.Next
		}

		// 移动到下一个节点
		head = head.Next
	}
	// 防止larger链表出现环，断开最后一个节点
	large.Next = nil
	// 连接两个链表
	small.Next = largeHead.Next

	// 返回新的链表头
	return smallHead.Next
}
