package swap_pairs

import "testing"

func TestSwapPairs(t *testing.T) {
	// 创建测试链表 1 -> 2 -> 3 -> 4
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}

	// 调用函数并打印结果
	newHead := swapPairs(head)
	printList(newHead) // 输出应为 2 1 4 3
}
