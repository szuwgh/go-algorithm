package add_two_numbers

import (
	"linklist"
)

// 题目描述：
// 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

// 示例：
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807

type LinkNode = linklist.LinkNode

func addTwoNumbers(l1 *LinkNode, l2 *LinkNode) *LinkNode {
	newNode := &LinkNode{}
	cur := newNode
	carry := 0 //进位
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Value
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Value
			l2 = l2.Next
		}
		//获取进位
		carry = sum / 10
		cur.Next = &linklist.LinkNode{Value: sum % 10}
		cur = cur.Next
	}
	return newNode.Next
}
