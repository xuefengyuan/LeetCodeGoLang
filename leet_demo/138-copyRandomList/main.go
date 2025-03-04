package main

import "fmt"

// 138 随机链表的复制
// 给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random ，该指针可以指向链表中的任何节点或空节点。
//
// 构造这个链表的 深拷贝。 深拷贝应该正好由 n 个 全新 节点组成，其中每个新节点的值都设为其对应的原节点的值。新节点的 next 指针和 random 指针也都应指向复制链表中的新节点，并使原链表和复制链表中的这些指针能够表示相同的链表状态。复制链表中的指针都不应指向原链表中的节点 。
//
// 例如，如果原链表中有 X 和 Y 两个节点，其中 X.random --> Y 。那么在复制链表中对应的两个节点 x 和 y ，同样有 x.random --> y 。
//
// 返回复制链表的头节点。
//
// 用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：
//
// val：一个表示 Node.val 的整数。
// random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为  null 。
// 你的代码 只 接受原链表的头节点 head 作为传入参数。
func main() {
	// 创建测试链表
	node1 := &Node{Val: 7}
	node2 := &Node{Val: 13}
	node3 := &Node{Val: 11}
	node4 := &Node{Val: 10}
	node5 := &Node{Val: 1}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	node1.Random = nil
	node2.Random = node1
	node3.Random = node5
	node4.Random = node3
	node5.Random = node1

	fmt.Println("原链表：")
	printList(node1)

	// 调用复制函数
	copyHead := copyRandomList(node1)

	fmt.Println("复制链表：")
	printList(copyHead)
}

// 打印链表，用于测试
func printList(head *Node) {
	curr := head
	for curr != nil {
		randomVal := "null"
		if curr.Random != nil {
			randomVal = fmt.Sprintf("%d", curr.Random.Val)
		}
		fmt.Printf("[Val: %d, Random: %s] -> ", curr.Val, randomVal)
		curr = curr.Next
	}
	fmt.Println("null")
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	// 处理空链表的情况
	if head == nil {
		return nil
	}

	// 第一步：拷贝每个节点并插入到当前节点后

	curr := head
	for curr != nil {
		copyNode := &Node{Val: curr.Val}
		copyNode.Next = curr.Next
		curr.Next = copyNode
		curr = copyNode.Next
	}

	// 第二步：设置拷贝节点的随机指针
	curr = head
	for curr != nil {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		}
		curr = curr.Next.Next
	}

	// 第三步：分离原链表和拷贝链表
	curr = head
	copyHead := head.Next
	for curr != nil {
		copyNode := curr.Next
		curr.Next = copyNode.Next
		if copyNode.Next != nil {
			copyNode.Next = copyNode.Next.Next
		}
		curr = curr.Next
	}

	return copyHead

}
