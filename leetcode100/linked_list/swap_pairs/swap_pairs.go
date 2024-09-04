package swap_pairs

import "fmt"

// 两两交换链表中的节点
// 中等
// 相关标签
// 相关企业
// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

// 示例 1：

// 1->2->3->4

// 2->1->4->3

// 输入：head = [1,2,3,4]
// 输出：[2,1,4,3]
// 示例 2：

// 输入：head = []
// 输出：[]
// 示例 3：

// 输入：head = [1]
// 输出：[1]

// 要解决这个问题，你可以使用 Go 语言编写一个
// 函数来两两交换链表中的节点。
// 这个问题可以通过迭代或递归来解决。下面是迭代方法的实现：

// 迭代方法
// 通过迭代的方式，我们可以一对一对地交换节点，直到链表末尾。
// 为了方便处理头节点的交换，我们可以使用一个虚拟节点（dummy node）。
// 代码说明
// ListNode 结构：定义了链表节点的数据结构。
// swapPairs 函数：负责两两交换链表中的节点。首先创建一个虚拟节点 dummy，并让 prev 指向 dummy。在循环中，每次提取当前的两个节点，然后进行交换，最后调整指针以继续处理后续节点。
// printList 函数：用于打印链表的值。
// main 函数：创建测试链表并调用 swapPairs 函数验证结果。
// 递归方法
// 也可以用递归的方法来实现，但递归方式通常在实际应用中不如迭代方式效率高。以上迭代方式是更常用的解决方案。

// ListNode 是链表的节点结构
type ListNode struct {
	Val  int
	Next *ListNode
}

// swapPairs 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	// 创建一个虚拟节点
	dummy := &ListNode{Next: head}
	prev := dummy

	// 遍历链表，交换每一对节点
	for head != nil && head.Next != nil {
		// 定义要交换的两个节点
		firstNode := head
		secondNode := head.Next

		// 交换节点
		prev.Next = secondNode
		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode

		// 移动到下一对节点
		prev = firstNode
		head = firstNode.Next
	}

	// 返回交换后的链表
	return dummy.Next
}

// 打印链表的工具函数
func printList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val, " ")
		node = node.Next
	}
	fmt.Println()
}
