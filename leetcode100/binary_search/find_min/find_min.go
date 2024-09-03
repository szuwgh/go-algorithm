package find_min

// 寻找旋转排序数组中的最小值

// 提示
// 已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,2,4,5,6,7] 在变化后可能得到：
// 若旋转 4 次，则可以得到 [4,5,6,7,0,1,2]
// 若旋转 7 次，则可以得到 [0,1,2,4,5,6,7]
// 注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。

// 给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。

// 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。

// 示例 1：

// 输入：nums = [3,4,5,1,2]
// 输出：1
// 解释：原数组为 [1,2,3,4,5] ，旋转 3 次得到输入数组。
// 示例 2：

// 输入：nums = [4,5,6,7,0,1,2]
// 输出：0
// 解释：原数组为 [0,1,2,4,5,6,7] ，旋转 3 次得到输入数组。
// 示例 3：

// 输入：nums = [11,13,15,17]
// 输出：11
// 解释：原数组为 [11,13,15,17] ，旋转 4 次得到输入数组。

// 要在旋转排序数组中找到最小值，我们可以使用二分搜索算法，这样可以在

// O(logn) 的时间复杂度内解决问题。下面是一个使用二分搜索的方法：

// 定义边界条件：
// 初始化左右边界 left 和 right，分别为数组的起始和结束索引。
// 二分搜索：
// 当 left 小于等于 right 时，计算中间索引 mid。
// 检查 nums[mid] 和 nums[right] 的关系：
// 如果 nums[mid] 大于 nums[right]，说明最小值在 mid 右侧，更新 left 为 mid + 1。
// 如果 nums[mid] 小于等于 nums[right]，说明最小值在 mid 左侧或就是 mid，更新 right 为 mid。
// 结束条件：
// 当 left 等于 right 时，left 或 right 就是最小值的索引，返回 nums[left] 即可。

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}
