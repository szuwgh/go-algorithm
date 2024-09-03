package main

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	// 示例链表：[1,2,3,4,5]
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}

	fmt.Print("Original List: ")
	printList(head)

	reversedHead := reverseList(head)
	fmt.Print("Reversed List: ")
	printList(reversedHead)
}
