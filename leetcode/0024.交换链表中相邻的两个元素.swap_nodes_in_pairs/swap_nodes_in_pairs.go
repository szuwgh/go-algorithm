package swap_nodes_in_pairs

import (
	"linklist"
)

// 交换链表中相邻的两个元素。
// 注意第一个节点与第二个节点要交换位置，而第二个节点不用与第三个节点交换位置。
// 注意点：
// 不允许修改节点的值
// 只能用常量的额外空间

//1->2->3->4, you should return the list as 2->1->4->3.

type linkNode = linklist.LinkNode

func swapPairs(head *linkNode) *linkNode {
	if head == nil || head.Next == nil {
		return head
	}

	newhead := head.Next
	head.Next = swapPairs(newhead.Next)
	newhead.Next = head
	return newhead
}
