package remove_nth_node_from_end_of_list

import "linklist"

// 题目描述：

// 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

// 示例：
// 给定一个链表: 1->2->3->4->5, 和 n = 2.
// 当删除了倒数第二个节点后，链表变为 1->2->3->5.
// 说明：
// 给定的 n 保证是有效的。
// 进阶：
// 你能尝试使用一趟扫描实现吗？

type LinkNode = linklist.LinkNode

func removeNthFromEnd(head *LinkNode, n int) *LinkNode {
	fast := head
	slow := head
	dumy := slow
	for i := 0; i <= n; i++ {
		if fast != nil {
			fast = fast.Next
		}
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dumy
}
