package reverse_list

import "fmt"

// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

// 示例 1：

// 输入：head = [1,2,3,4,5]
// 输出：[5,4,3,2,1]
// 示例 2：

// 输入：head = [1,2]
// 输出：[2,1]
// 示例 3：

// 输入：head = []
// 输出：[]

// 初始化三个指针：prev、curr 和 next。prev 指针用来保存反转后的链表的前一个节点，curr 指针用来遍历原链表，next 指针用来保存当前节点的下一个节点。
// 遍历原链表，将每个节点的 next 指针指向其前一个节点，完成链表的反转。
// 更新 prev 和 curr 指针以继续遍历。
// 最终返回 prev，它指向反转后的链表的头节点

// 在这个示例中：

// ListNode 结构体定义了链表的节点。
// reverseList 函数接受链表头节点并返回反转后的链表头节点。
// printList 函数用于打印链表中的值。
// main 函数创建一个示例链表，调用 reverseList 函数来反转链表，并打印原始链表和反转后的链表。

// 定义链表节点结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转链表函数
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next // 保存当前节点的下一个节点
		curr.Next = prev  // 反转当前节点的指针
		prev = curr       // 更新 prev 为当前节点
		curr = next       // 更新 curr 为下一个节点
	}

	return prev // prev 现在指向新的头节点
}

// 打印链表函数
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}
