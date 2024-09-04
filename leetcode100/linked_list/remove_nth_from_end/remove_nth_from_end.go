package removeNthFromEnd

import (
	"fmt"
)

// 删除链表的倒数第 N 个结点
// 中等
// 相关标签
// 相关企业
// 提示
// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

// 示例 1：

// 1 -> 2 -> 3 -> 4 -> 5
// 1 -> 2 -> 3 ---> 5

// 输入：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]
// 示例 2：

// 输入：head = [1], n = 1
// 输出：[]
// 示例 3：

// 输入：head = [1,2], n = 1
// 输出：[1]

// 要删除链表的倒数第 n 个节点，
// 可以使用双指针的方法。首先让两个指针之间保持 n 的距离，
// 然后同时移动两个指针，当第一个指针到达链表末尾时，
// 第二个指针的位置就是要删除节点的前一个节点。
// 接下来删除该节点，并返回链表的头节点。

// 代码说明
// 哑节点：dummy 节点作为头节点的前驱节点，方便处理删除头节点的情况。
// 双指针：first 和 second 两个指针，初始都指向 dummy 节点。首先移动 first 指针 n+1 步，
// 然后同时移动 first 和 second 指针，直到 first 到达末尾。
// 删除节点：此时 second 的下一个节点就是要删除的节点，执行删除操作。

// ListNode 定义链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// removeNthFromEnd 删除链表的倒数第 n 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 创建一个哑节点指向头节点，以应对删除头节点的情况
	dummy := &ListNode{Next: head}
	first := dummy
	second := dummy

	// 先让第一个指针移动n+1步
	for i := 0; i <= n; i++ {
		first = first.Next
	}

	// 然后同时移动两个指针，直到第一个指针到达末尾
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// 此时第二个指针的下一个节点就是需要删除的节点
	second.Next = second.Next.Next

	// 返回链表头节点
	return dummy.Next
}

// 打印链表
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}
