package removeNthFromEnd

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
	// 构造测试链表: [1,2,3,4,5]
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}

	// 删除倒数第2个结点
	newHead := removeNthFromEnd(head, 2)

	// 打印结果链表
	printList(newHead)
}
