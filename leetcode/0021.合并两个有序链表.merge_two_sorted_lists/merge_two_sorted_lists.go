package merge_two_sorted_lists

import "linklist"

// 21.合并两个有序链表
// 题目描述：
// 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

// 示例：
// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4

type LinkNode = linklist.LinkNode

func mergeTwoSortedList(l1 *LinkNode, l2 *LinkNode) *LinkNode {
	dummy := &LinkNode{}
	temp := dummy
	for l1 != nil && l2 != nil {
		if l1.Value < l2.Value {
			temp.Next = l1
			l1 = l1.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
		}
		temp = temp.Next
	}
	if l1 != nil {
		temp.Next = l1
	} else if l2 != nil {
		temp.Next = l2
	}
	return dummy.Next
}
