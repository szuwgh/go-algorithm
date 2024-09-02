package sort_list

import "linklist"

//链表排序

type linkNode = linklist.LinkNode

//排序
func sortList(head *linkNode) *linkNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := getMid(head)
	second := mid.Next
	mid.Next = nil
	return merge(sortList(head), sortList(second))
}

//归并排序
func merge(first, second *linkNode) *linkNode {
	dummy := &linkNode{}
	pre := dummy
	for first != nil && second != nil {
		if first.Value > second.Value {
			pre.Next = second
			second = second.Next
		} else {
			pre.Next = first
			first = first.Next
		}
		pre = pre.Next
	}
	if first == nil {
		pre.Next = second
	} else {
		pre.Next = first
	}
	return dummy.Next
}

//获取中间节点
//快慢指针
func getMid(head *linkNode) *linkNode {
	fast := head
	slow := head
	for fast.Next.Next != nil && slow.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
