package longestconsecutive

// 最长连续序列

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
// 示例 1：

// 输入：nums = [100,4,200,1,3,2]
// 输出：4
// // 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
// 示例 2：

// 输入：nums = [0,3,7,2,5,8,4,6,0,1]
// 输出：9

func longestconsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 将所有数字加入哈希表
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	longest := 0

	for _, num := range nums {
		// 只有当这个数字的前一个数字不在哈希表中时，才开始新的序列
		if !numSet[num-1] {
			currnetNum := num
			currnetStreak := 1

			for numSet[currnetNum+1] {
				currnetNum++
				currnetStreak++
			}
			if currnetStreak > longest {
				longest = currnetStreak
			}

		}
	}
	return longest
}
