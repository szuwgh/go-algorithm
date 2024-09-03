package rob

// 打家劫舍
// 你是一个专业的小偷，计划偷窃沿街的房屋。
// 每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
// 如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
// 给定一个代表每个房屋存放金额的非负整数数组，
// 计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

// 这个问题是经典的动态规划问题，称为“打家劫舍”问题。我们可以通过动态规划来解决这个问题。

// 问题描述：
// 给定一个非负整数数组，每个元素代表一间房屋中存放的现金金额，要求找出在不偷相邻房屋的前提下，能够偷窃的最大金额。

// 解决思路：
// 用一个动态规划数组 dp 来记录到达每个房屋时能够偷窃的最大金额。

// 对于第 i 个房屋，有两个选择：

// 不偷当前房屋，最大金额就是 dp[i-1]。
// 偷当前房屋，那么最大金额就是当前房屋的金额加上前两个房屋的最大金额，即 nums[i] + dp[i-2]。
// 转移方程：dp[i] = max(dp[i-1], nums[i] + (dp[i-2] if i >= 2 else 0))

// 初始条件：

// dp[0] = nums[0]
// dp[1] = max(nums[0], nums[1])
// 最终，dp[n-1] 就是我们要求的最大金额，其中 n 是房屋的数量。

// 示例 1：

// 输入：[1,2,3,1]
// 输出：4
// 解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
//      偷窃到的最高金额 = 1 + 3 = 4 。
// 示例 2：

// 输入：[2,7,9,3,1]
// 输出：12
// 解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
//      偷窃到的最高金额 = 2 + 9 + 1 = 12 。
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
