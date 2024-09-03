package merge_two_list

import (
	"fmt"
)

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

// 示例 1：

// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
// 示例 2：

// 输入：l1 = [], l2 = []
// 输出：[]
// 示例 3：

// 输入：l1 = [], l2 = [0]
// 输出：[0]

// 解释：
// ListNode 结构体: 定义了链表节点的结构，每个节点包含一个整数值 Val 和一个指向下一个节点的指针 Next。

// mergeTwoLists 函数:

// 该函数接受两个链表的头节点 l1 和 l2。
// 它创建了一个哨兵节点 dummy，以及一个指针 tail，用于跟踪新链表的末尾。
// 在两个链表都不为空时，比较两个链表当前节点的值，将较小的节点链接到新链表中，并移动相应链表的指针。
// 当一个链表耗尽后，将另一个链表的剩余部分直接链接到新链表末尾。
// printList 函数: 用于打印链表中的所有节点值。

// main 函数: 测试了三个不同的例子，展示了如何使用 mergeTwoLists 函数。

// 这样就可以成功地将两个升序链表合并为一个新的升序链表。

// 定义链表节点结构
type ListNode struct {
	Val  int
	Next *ListNode
}

// 合并两个升序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 创建一个哨兵节点，方便处理链表的合并
	dummy := &ListNode{}
	// tail 指针用于构建新链表
	tail := dummy

	// 当 l1 和 l2 都不为空时，依次比较每个节点的值
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	// 如果 l1 还有剩余节点，直接接到新链表的尾部
	if l1 != nil {
		tail.Next = l1
	}

	// 如果 l2 还有剩余节点，直接接到新链表的尾部
	if l2 != nil {
		tail.Next = l2
	}

	// 返回哨兵节点的下一个节点，即合并后的链表头部
	return dummy.Next
}

// 辅助函数，用于打印链表
func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d ", node.Val)
		node = node.Next
	}
	fmt.Println()
}
