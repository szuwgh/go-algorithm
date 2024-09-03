package search_insert

// 搜索插入位置
// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

// 请必须使用时间复杂度为 O(log n) 的算法。

// 示例 1:

// 输入: nums = [1,3,5,6], target = 5
// 输出: 2
// 示例 2:

// 输入: nums = [1,3,5,6], target = 2
// 输出: 1
// 示例 3:

// 输入: nums = [1,3,5,6], target = 7
// 输出: 4
// 代码说明
// 变量初始化: left 初始化为数组的起始位置 0，right 初始化为数组的长度 len(nums)。

// 二分查找:

// 计算中间索引 mid。
// 如果 nums[mid] 等于目标值 target，直接返回 mid。
// 如果 nums[mid] 小于目标值 target，说明目标值在右半部分，将 left 更新为 mid + 1。
// 如果 nums[mid] 大于目标值 target，说明目标值在左半部分，将 right 更新为 mid。
// 返回结果: 如果没有找到目标值，left 将会是目标值应该插入的位置。

// 这个实现保证了时间复杂度为

// O(logn)，符合要求。
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)

	// 使用二分查找法
	for left < right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
