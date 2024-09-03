package main

// 环形链表
// 简单
// 相关标签
// 相关企业
// 给你一个链表的头节点 head ，判断链表中是否有环。

// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。

// 如果链表中存在环 ，则返回 true 。 否则，返回 false 。

// 示例 1：

// 输入：head = [3,2,0,-4], pos = 1
// 输出：true
// 解释：链表中有一个环，其尾部连接到第二个节点。
// 示例 2：

// 输入：head = [1,2], pos = 0
// 输出：true
// 解释：链表中有一个环，其尾部连接到第一个节点。
// 示例 3：

// 要判断一个链表是否存在环，可以使用快慢指针（Floyd 判圈算法）。
// 该算法使用两个指针，一个快指针（fast）和一个慢指针（slow）。
// 慢指针一次走一步，快指针一次走两步。
// 如果链表中存在环，则快指针最终会追上慢指针，说明链表中存在环；
// 如果快指针走到链表的末尾（即 fast 或 fast.next 为 nil），说明链表中不存在环

// 代码说明：
// ListNode 结构体：定义了链表的节点，每个节点包含一个整数值和一个指向下一个节点的指针。
// hasCycle 函数：判断链表中是否存在环。首先检查头节点是否为空或头节点的下一个节点是否为空，如果是，则直接返回 false。否则，使用快慢指针方法来检测环。
// main 函数：创建一个示例链表 [3,2,0,-4] 并在 pos = 1 处形成环，然后调用 hasCycle 函数来判断链表是否有环，并输出结果。

// ListNode 定义链表的节点结构
type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle 判断链表中是否存在环
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}
