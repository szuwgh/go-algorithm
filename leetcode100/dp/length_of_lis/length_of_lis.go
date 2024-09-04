package lengthOfLIS

// 最长递增子序列

// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的
// 子序列
// 。

// 示例 1：

// 输入：nums = [10,9,2,5,3,7,101,18]
// 输出：4
// 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
// 示例 2：

// 输入：nums = [0,1,0,3,2,3]
// 输出：4
// 示例 3：

// 输入：nums = [7,7,7,7,7,7,7]
// 输出：1

// 代码解释：
// dp[i] 表示以 nums[i] 结尾的最长递增子序列的长度。
// 外层循环 i 遍历数组 nums，内层循环 j 用于检查 nums[i] 之前的元素是否小于 nums[i]。
// 如果 nums[i] > nums[j]，说明 nums[i] 可以接在 nums[j] 后面形成递增子序列，这时更新 dp[i]。
// 最后，通过 maxLength 变量保存全局的最长递增子序列的长度。
// 示例运行结果：
// 对于 nums = [10, 9, 2, 5, 3, 7, 101, 18]，最长递增子序列是 [2, 3, 7, 101]，长度为 4。
// 对于 nums = [0, 1, 0, 3, 2, 3]，最长递增子序列长度为 4。
// 对于 nums = [7, 7, 7, 7, 7, 7, 7]，最长递增子序列长度为 1。
// 这种方法的时间复杂度是 O(n^2)，其中 n 是数组的长度。

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// dp[i] 表示以 nums[i] 结尾的最长递增子序列的长度
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1 // 每个单独的元素本身就是长度为 1 的递增子序列
	}

	maxLength := 1 // 最长递增子序列的长度至少为 1

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				// 如果 nums[i] 能够接在 nums[j] 后面形成递增序列，则更新 dp[i]
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLength = max(maxLength, dp[i]) // 更新全局的最长递增子序列长度
	}

	return maxLength
}

// 辅助函数，返回两个整数中的较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
