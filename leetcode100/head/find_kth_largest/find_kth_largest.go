package find_kth_larget

import (
	"math/rand"
	"time"
)

// 数组中的第K个最大元素
// 中等
// 相关标签
// 相关企业
// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

// 示例 1:

// 输入: [3,2,1,5,6,4], k = 2
// 输出: 5
// 示例 2:

// 输入: [3,2,3,1,2,4,5,5,6], k = 4
// 输出: 4

// 提示：

// 1 <= k <= nums.length <= 105
// -104 <= nums[i] <= 104

// 要在 Golang 中找到数组的第 k 大元素，且时间复杂度要求为 O(n)，可以使用快速选择算法（Quickselect）。
// 这种算法与快速排序（Quicksort）类似，但它只递归处理一半的数据，从而将平均时间复杂度降低到 O(n)。

// 快速选择算法简介
// 快速选择算法的核心思想与快速排序相同。它基于分区的思想，每次选择一个基准元素（pivot），将数组分成两部分：小于基准的部分和大于基准的部分。
// 如果分区点正好是第 k 大的位置，就返回该元素，否则根据分区结果选择继续在左半部分或右半部分递归查找。

// 代码说明
// partition 函数通过随机选择一个 pivot，将数组分为两部分：大于 pivot 的元素放在左边，小于 pivot 的元素放在右边。最终返回 pivot 的正确位置。
// quickSelect 函数递归地选择哪一部分继续进行分区，直到找到第 k 大的元素。
// findKthLargest 函数是主函数，调用 quickSelect 并传入 k-1 作为索引（因为索引从 0 开始）。
// 复杂度分析
// 时间复杂度：平均时间复杂度为 O(n)，最坏情况下可能达到 O(n²)（如果每次选择的 pivot 都非常不平衡），但通过随机化 pivot 的选择，可以降低最坏情况发生的概率。
// 空间复杂度：O(1)（不考虑递归栈的空间消耗）。

// partition 函数用于分区，返回 pivot 的最终位置
func partition(nums []int, left, right int) int {
	// 随机选择一个 pivot 元素
	rand.Seed(time.Now().UnixNano())
	pivotIndex := rand.Intn(right-left+1) + left
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]

	pivot := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] > pivot { // 大于 pivot 的元素放在前面
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	// 将 pivot 放到它的最终位置
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

// quickSelect 函数用于找到第 k 大的元素
func quickSelect(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}

	// 分区
	pivotIndex := partition(nums, left, right)

	// 根据 pivotIndex 选择递归方向
	if pivotIndex == k {
		return nums[pivotIndex]
	} else if pivotIndex > k {
		return quickSelect(nums, left, pivotIndex-1, k)
	} else {
		return quickSelect(nums, pivotIndex+1, right, k)
	}
}

// findKthLargest 找到数组中第 k 大的元素
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k-1)
}
