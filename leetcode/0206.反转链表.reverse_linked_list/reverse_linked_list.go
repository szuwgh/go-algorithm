package reverse_linked_list

import "linklist"

// 题目描述：
// 反转一个单链表。

// 示例:
// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL

type linkNode = linklist.LinkNode

func reverseLinkedList(head *linkNode) *linkNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *linkNode
	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}

func reverseLinkedList2(head *linkNode) *linkNode {
	if head == nil || head.Next == nil {
		return head
	}
	newhead := reverseLinkedList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newhead
}
