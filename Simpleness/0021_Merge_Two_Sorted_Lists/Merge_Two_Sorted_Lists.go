/**
 * @Author: darry
 * @Desc: 合并两个有序链表
Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
 * @Date: 2024/3/13 22:14
*/
package _021_Merge_Two_Sorted_Lists

import "github.com.leetCodeGo/structures"

type ListNode = structures.ListNode

//Definition for singly-linked list.
//type ListNode struct {
//    Val  int
//    Next *ListNode
//}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }

    // 链表1的值小于链表2值，进行递归，递归要传入一下个链表的node
    if l1.Val < l2.Val {
        l1.Next = mergeTwoLists(l1.Next, l2)
        return l1
    }
    // 递归，链表2要传入下一个node
    l2.Next = mergeTwoLists(l1, l2.Next)

    return l2
}
