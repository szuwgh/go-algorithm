package merge_two_list

import "testing"

func TestMergeTwoList(t *testing.T) {
	// 示例 1: l1 = [1, 2, 4], l2 = [1, 3, 4]
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	printList(l1)
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	printList(l2)
	mergedList := mergeTwoLists(l1, l2)
	printList(mergedList) // 输出: 1 1 2 3 4 4

	// 示例 2: l1 = [], l2 = []
	l1 = nil
	l2 = nil
	mergedList = mergeTwoLists(l1, l2)
	printList(mergedList) // 输出: (空)

	// 示例 3: l1 = [], l2 = [0]
	l1 = nil
	l2 = &ListNode{0, nil}
	mergedList = mergeTwoLists(l1, l2)
	printList(mergedList) // 输出: 0
}
