package main

// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。

// 子数组是数组中元素的连续非空序列。

// 示例 1：

// 输入：nums = [1,1,1], k = 2
// 输出：2
// 示例 2：

// 输入：nums = [1,2,3], k = 3
// 输出：2

// 要解决这个问题，我们可以使用前缀和（Prefix Sum）加上哈希表的方法来高效地计算和为 k 的子数组的个数。具体思路如下：

// 前缀和的概念：前缀和是指数组中从第一个元素到当前元素的和。例如，对于数组 [1, 2, 3]，它的前缀和依次为 [1, 3, 6]。

// 计算前缀和：假设 sum(i) 表示数组 nums 从起始位置到第 i 个元素的和，那么子数组 nums[i:j] 的和 sum(i,j) 可以表示为 sum(j) - sum(i-1)。

// 哈希表存储前缀和：我们使用一个哈希表 prefixSumCount 来存储前缀和出现的次数。遍历数组时，计算当前前缀和 currentSum，并检查 currentSum - k 是否存在于哈希表中。如果存在，则表示存在一个子数组的和为 k。最后，把当前的前缀和 currentSum 加入哈希表中。

// 初始条件：为了处理数组从开始部分就满足和为 k 的情况，我们初始化 prefixSumCount[0] = 1。

func subarraySum(nums []int, k int) int {
	prefixSumCount := make(map[int]int)
	prefixSumCount[0] = 1 // 初始化前缀和为0的出现次数为1

	currentSum := 0
	count := 0

	for _, num := range nums {
		currentSum += num

		if val, ok := prefixSumCount[currentSum-k]; ok {
			count += val
		}

		prefixSumCount[currentSum]++
	}

	return count
}
