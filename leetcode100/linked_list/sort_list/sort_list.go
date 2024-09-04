package sort_list

// 两两交换链表中的节点
// 中等
// 相关标签
// 相关企业
// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

// 示例 1：

// 输入：head = [1,2,3,4]
// 输出：[2,1,4,3]
// 示例 2：

// 输入：head = []
// 输出：[]
// 示例 3：

// 输入：head = [1]
// 输出：[1]

// 代码说明
// sortList 函数：

// 通过快慢指针找到链表的中点，将链表分成两部分。
// 递归地对这两部分链表进行排序。
// 最后合并两个已排序的链表。
// merge 函数：

// 合并两个有序链表，返回合并后的有序链表。
// arrayToList 和 listToArray 函数：

// 用于测试：arrayToList 将数组转换为链表，listToArray 将链表转换为数组，方便打印结果。
// 测试用例
// 输入 [4, 2, 1, 3]，输出 [1, 2, 3, 4]。
// 输入 [-1, 5, 3, 4, 0]，输出 [-1, 0, 3, 4, 5]。
// 输入 []，输出 []。

type ListNode struct {
	Val  int
	Next *ListNode
}

// sortList 对链表进行排序
func sortList(head *ListNode) *ListNode {
	// 如果链表为空或者只有一个节点，直接返回
	if head == nil || head.Next == nil {
		return head
	}

	// 使用快慢指针找到链表中间的节点
	slow, fast := head, head
	var prev *ListNode

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 将链表分为两部分
	prev.Next = nil

	// 递归对两部分链表排序
	left := sortList(head)
	right := sortList(slow)

	// 合并排序后的两部分链表
	return merge(left, right)
}

// merge 合并两个已排序的链表
func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}

	return dummy.Next
}

// 辅助函数：将数组转换为链表
func arrayToList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	current := head
	for _, val := range arr[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

// 辅助函数：将链表转换为数组
func listToArray(head *ListNode) []int {
	var arr []int
	current := head
	for current != nil {
		arr = append(arr, current.Val)
		current = current.Next
	}
	return arr
}
