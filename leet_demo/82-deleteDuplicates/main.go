package main

import "fmt"

// 82 删除排序链表中的重复元素II
// 给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字。返回已排序的链表。
func main() {
	// 创建测试链表: 1 -> 2 -> 3 -> 3 -> 4 -> 4 -> 5
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 5}

	fmt.Println("原链表:")
	printList(head)

	// 删除重复节点
	newHead := deleteDuplicates(head)

	fmt.Println("删除重复节点后的链表:")
	printList(newHead)
}

// 打印链表
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

func deleteDuplicates(head *ListNode) *ListNode {
	// 创建一个哑节点，方便处理头节点可能被删除的情况
	dummy := &ListNode{Next: head}
	// 定义一个指针，用于遍历链表
	prve := dummy
	for head != nil {
		// 如果下一节点不为空，并且当前节点和下一个节点值相同
		if head.Next != nil && head.Val == head.Next.Val {
			// 循环跳过所有值相同的节点
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			// 将 prev 的 Next 指向非重复元素
			prve.Next = head.Next
		} else {
			// 如果当前节点不重复，移动 prev 指针
			prve = prve.Next
		}

		// 移动 head 指针
		head = head.Next
	}
	// 返回去重后的链表
	return dummy.Next
}
