package main

import (
	"fmt"
	"testing"
)

func TestHasCycle(t *testing.T) {
	// 创建一个示例链表 [3,2,0,-4] 并在 pos = 1 处形成环
	node1 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 0}
	node4 := &ListNode{Val: -4}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2 // 形成环

	result := hasCycle(node1)
	fmt.Println("是否存在环:", result) // 输出 true
}
