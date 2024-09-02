package reverse_nodes_in_k_group

import "linklist"

//将一个链表中每k个数进行翻转，末尾不足k个的数不做变化。
// Example:

// Given this linked list: 1->2->3->4->5

// For k = 2, you should return: 2->1->4->3->5

// For k = 3, you should return: 3->2->1->4->5
type linkNode = linklist.LinkNode

func reverseNodesInKGroup(head *linkNode, k int) *linkNode {
	if head == nil || head.Next == nil {
		return head
	}
	tail, ok := getTail(head, k)
	if ok {
		tailNext := tail.Next
		tail.Next = nil
		head, tail = reverse(head, tail)
		/* tail 后面接上尾部的递归处理 */
		tail.Next = reverseNodesInKGroup(tailNext, k)
	}
	return head
}

func getTail(head *linkNode, k int) (*linkNode, bool) {
	for k > 1 && head != nil {
		head = head.Next
		k--
	}
	return head, k == 1 && head != nil
}

func reverse(head, tail *linkNode) (*linkNode, *linkNode) {
	var prev *linkNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	return prev, head
}
