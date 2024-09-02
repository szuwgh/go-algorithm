package rotate_List

import "linklist"

// Input: 1->2->3->4->5->NULL, k = 2
// Output: 4->5->1->2->3->NULL
// Explanation:
// rotate 1 steps to the right: 5->1->2->3->4->NULL
// rotate 2 steps to the right: 4->5->1->2->3->NULL

type linkNode = linklist.LinkNode

func rotateList(head *linkNode, k int) *linkNode {
	if head == nil {
		return nil
	}
	n := 1
	cur := head
	for cur.Next != nil {
		n++
		cur = cur.Next
	}
	cur.Next = head
	m := n - k%n
	for i := 0; i < m; i++ {
		cur = cur.Next
	}
	newHead := cur.Next
	cur.Next = nil
	return newHead
}
