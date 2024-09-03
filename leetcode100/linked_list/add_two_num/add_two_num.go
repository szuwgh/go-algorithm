package addtwonum

import "fmt"

// 两数相加

// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

// 请你将两个数相加，并以相同形式返回一个表示和的链表。

// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

// 示例 1：

// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.
// 示例 2：

// 输入：l1 = [0], l2 = [0]
// 输出：[0]
// 示例 3：

// 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// 输出：[8,9,9,9,0,0,0,1]

// 解释：
// ListNode 结构体：定义了一个链表节点，包含一个整数 Val 和一个指向下一个节点的指针 Next。

// addTwoNumbers 函数：

// 使用一个哑节点 dummy，方便后续操作。
// 使用 carry 变量来处理每一位相加后的进位。
// 遍历两个链表，将对应位置的值相加，并创建新节点保存相加结果的个位部分。
// 如果有进位，需要在链表的最后添加一个新节点。
// 辅助函数：

// createList 函数根据给定的整数切片创建对应的链表。
// printList 函数用于打印链表。

// 定义链表节点结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

// 两数相加函数
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 创建一个哑节点(dummy node)
	dummy := &ListNode{}
	current := dummy
	carry := 0

	// 遍历两个链表直到两个链表都为空
	for l1 != nil || l2 != nil {
		// 获取当前节点的值，如果链表为空则取0
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		// 计算当前位的和以及进位
		sum := carry + x + y
		carry = sum / 10

		// 将结果添加到新链表中
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	// 如果最终有进位，需要在链表最后添加一个节点
	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}

	// 返回结果链表的头节点（哑节点的下一个节点）
	return dummy.Next
}

// 辅助函数：根据切片创建链表
func createList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

// 辅助函数：打印链表
func printList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val)
		if node.Next != nil {
			fmt.Print(" -> ")
		}
		node = node.Next
	}
	fmt.Println()
}
