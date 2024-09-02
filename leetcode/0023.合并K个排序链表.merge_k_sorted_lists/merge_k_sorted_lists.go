package merge_k_sorted_lists

import "linklist"

// 到题目第个想到的是归并排序过程中的归并操作子过程，从头开始两两比较，
// 找出最小的，然后接着往后比较，常用的是2路归并。而题目给的是k个已排好序的链表(k>=2)。
// 如果没有提示，我半天不知道如何去实现，幸好提示说用最小堆来做k路合并，
// 于是我想到可以这样做：创建一个大小为k的数组，将k个链表中的第一个元素依次存放到数组中，
// 然后将数组调整为最小堆，这样保证数组的第一个元素是最小的，假设为min，将min从最小堆
// 取出并存放到最终结果的链表中，此时将min所在链表的下一个元素到插入的最小堆中，继续上面的
// 操作，直到堆中没有元素为止
// 输入:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// 输出: 1->1->2->3->4->4->5->6

type LinkNode = linklist.LinkNode

func mergeKSortsLists(ln []*LinkNode) *LinkNode {
	size := len(ln)
	listNode := make([]*LinkNode, size)

	for i := 0; i < size; i++ {
		listNode[i] = ln[i]
	}
	return nil
}
