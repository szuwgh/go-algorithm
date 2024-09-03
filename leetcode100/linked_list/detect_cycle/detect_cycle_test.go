package detect_cycle

import (
	"fmt"
	"testing"
)

// createLinkedListWithCycle 创建一个带环的链表
func createLinkedListWithCycle(values []int, pos int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head
	var cycleEntry *ListNode

	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
		if i == pos {
			cycleEntry = current
		}
	}

	if pos >= 0 {
		current.Next = cycleEntry
	}

	return head
}

func TestDetectCycle(t *testing.T) {
	// 示例 1
	values1 := []int{3, 2, 0, -4}
	pos1 := 1
	head1 := createLinkedListWithCycle(values1, pos1)
	cycleNode1 := detectCycle(head1)
	if cycleNode1 != nil {
		fmt.Printf("示例 1：链表中环的起始节点的值是: %d\n", cycleNode1.Val)
	} else {
		fmt.Println("示例 1：链表中没有环")
	}

	// 示例 2
	values2 := []int{1, 2}
	pos2 := 0
	head2 := createLinkedListWithCycle(values2, pos2)
	cycleNode2 := detectCycle(head2)
	if cycleNode2 != nil {
		fmt.Printf("示例 2：链表中环的起始节点的值是: %d\n", cycleNode2.Val)
	} else {
		fmt.Println("示例 2：链表中没有环")
	}

	// 示例 3
	values3 := []int{1}
	pos3 := -1
	head3 := createLinkedListWithCycle(values3, pos3)
	cycleNode3 := detectCycle(head3)
	if cycleNode3 != nil {
		fmt.Printf("示例 3：链表中环的起始节点的值是: %d\n", cycleNode3.Val)
	} else {
		fmt.Println("示例 3：链表中没有环")
	}
}
