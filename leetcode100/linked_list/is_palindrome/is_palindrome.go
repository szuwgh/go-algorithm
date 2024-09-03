package main

import "fmt"

// 回文链表

// 给你一个单链表的头节点 head ，请你判断该链表是否为
// 回文链表
// 。如果是，返回 true ；否则，返回 false 。

// 示例 1：

// 输入：head = [1,2,2,1]
// 输出：true
// 示例 2：

// 输入：head = [1,2]

// 找到链表的中间节点：使用快慢指针，快指针每次移动两步，慢指针每次移动一步。当快指针到达链表末尾时，慢指针刚好到达链表的中间位置。

// 反转链表的后半部分：从中间位置开始，反转链表的后半部分。

// 比较链表的前半部分和反转后的后半部分：同时从链表的起始位置和反转后的中间位置进行比较。如果所有对应的节点值都相等，则链表是回文的。
// 解释
// ListNode 结构：定义了一个链表节点的结构体，包含一个整数值 Val 和指向下一个节点的指针 Next。
// isPalindrome 函数：判断链表是否为回文。首先通过快慢指针找到中间节点，然后反转后半部分的链表，最后逐个比较前半部分和反转后的后半部分。
// 主函数 main：创建两个示例链表并调用 isPalindrome 函数进行测试。

type ListNode struct {
	Val  int
	Next *ListNode
}

// 打印链表
func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println() // 打印换行符
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 步骤 1: 使用快慢指针找到链表的中间位置
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 步骤 2: 反转链表的后半部分
	var prev *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	// 步骤 3: 比较前半部分和反转后的后半部分
	left, right := head, prev
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}

	return true
}
