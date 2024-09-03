package jump

// 跳跃游戏 II
// 中等

// 给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。

// 每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:

// 0 <= j <= nums[i]
// i + j < n
// 返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。

// 示例 1:

// 输入: nums = [2,3,1,1,4]
// 输出: 2
// 解释: 跳到最后一个位置的最小跳跃数是 2。
//      从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
// 示例 2:

// 输入: nums = [2,3,0,1,4]
// 输出: 2

// 跳跃游戏 II 的问题可以用贪心算法来解决。我们需要找到到达数组最后一个位置的最小跳跃次数。我们可以通过以下步骤来实现这一目标：

// 初始化：

// jumps：跳跃次数，初始为0。
// current_end：当前跳跃的最远位置，初始为0。
// farthest：从当前跳跃位置能够到达的最远位置，初始为0。
// 遍历数组：

// 遍历到数组的每一个元素 i。
// 更新 farthest 为当前跳跃位置能够到达的最远位置。
// 如果 i 到达了 current_end，则更新 current_end 为 farthest 并且增加跳跃次数 jumps。
// 结束条件：

// 如果我们已经能够到达或超过了数组的最后一个位置，则返回跳跃次数 jumps。

// 解释：
// farthest：记录当前能到达的最远位置。
// currentEnd：表示当前跳跃能到达的最远位置。
// 当遍历到 currentEnd 的时候，我们增加跳跃次数，并更新 currentEnd 为 farthest，以准备下一次跳跃。
// 当 i 达到 n-1 之前，已经不需要额外的跳跃次数。
// 这样，我们通过贪心策略，确保每一步都尽可能跳得更远，从而计算出到达最后位置的最小跳跃次数。

func jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	jumps := 0
	currentEnd := 0
	farthest := 0

	for i := 0; i < n-1; i++ {
		farthest = max(farthest, i+nums[i])
		if i == currentEnd {
			jumps++
			currentEnd = farthest
		}
	}

	return jumps
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
