package topK

import "container/heap"

// golang 前 K 个高频元素
// 中等
// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

// 示例 1:

// 输入: nums = [1,1,1,2,2,3], k = 2
// 输出: [1,2]
// 示例 2:

// 输入: nums = [1], k = 1
// 输出: [1]

// 提示：

// 1 <= nums.length <= 105
// k 的取值范围是 [1, 数组中不相同的元素的个数]
// 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的

// 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。

// 在这道题目中，给定一个整数数组 nums 和一个整数 k，需要返回其中出现频率前 k 高的元素。
// 可以使用哈希表和堆来解决这个问题，并且我们要确保算法的时间复杂度优于

// 解题思路：
// 统计频率：我们可以使用哈希表来统计每个元素的出现次数，时间复杂度为

// 维护前 K 个高频元素：为了找到频率最高的前 k 个元素，可以使用小顶堆。堆的大小固定为 k，
// 这样我们可以确保堆中始终保存的是前 k 个频率最高的元素。

// 遍历哈希表的过程中，当堆的大小小于 k 时，直接将元素加入堆。
// 当堆的大小等于 k 且当前元素的频率大于堆顶元素时，弹出堆顶元素，将当前元素加入堆。

// 定义一个元素的结构体，用来存储元素和它的频率
type Element struct {
	value int
	count int
}

// 实现小顶堆
type MinHeap []Element

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].count < h[j].count }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	// 1. 统计每个元素的频率
	frequencyMap := make(map[int]int)
	for _, num := range nums {
		frequencyMap[num]++
	}

	// 2. 使用小顶堆来维护前 k 个高频元素
	h := &MinHeap{}
	heap.Init(h)

	for num, count := range frequencyMap {
		heap.Push(h, Element{value: num, count: count})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// 3. 将堆中的元素输出
	result := make([]int, 0, k)
	for h.Len() > 0 {
		result = append(result, heap.Pop(h).(Element).value)
	}

	return result
}
