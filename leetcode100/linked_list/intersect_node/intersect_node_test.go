package intersect_node

import (
	"fmt"
	"testing"
)

func TestIntersectNodego(t *testing.T) {
	intersect := &ListNode{Val: 2}
	intersect.Next = &ListNode{Val: 4}
	headA := &ListNode{Val: 1}
	headA.Next = &ListNode{Val: 9}
	headA.Next.Next = &ListNode{Val: 1}
	headA.Next.Next.Next = intersect

	// 创建链表 B
	headB := &ListNode{Val: 3}
	headB.Next = intersect

	// 打印链表
	fmt.Println("链表 A:")
	printList(headA)
	fmt.Println("链表 B:")
	printList(headB)

	// 查找相交节点
	intersection := getIntersectionNode(headA, headB)

	// 打印相交节点的值
	if intersection != nil {
		fmt.Printf("相交节点的值: %d\n", intersection.Val)
	} else {
		fmt.Println("两个链表没有相交")
	}

}
