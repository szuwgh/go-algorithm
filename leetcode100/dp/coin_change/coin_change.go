package coin_change

import "math"

// 零钱兑换
// 中等
// 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

// 你可以认为每种硬币的数量是无限的。

// 示例 1：

// 输入：coins = [1, 2, 5], amount = 11
// 输出：3
// 解释：11 = 5 + 5 + 1
// 示例 2：

// 输入：coins = [2], amount = 3
// 输出：-1
// 示例 3：

// 输入：coins = [1], amount = 0
// 输出：0

// 这是一个典型的动态规划问题，可以通过构建一个数组 dp 来解决这个问题。dp[i] 表示组成金额 i 所需的最少硬币数。

// 思路
// 初始化: 创建一个大小为 amount + 1 的数组 dp，将所有元素初始化为一个非常大的数值
// （例如 amount + 1），因为我们需要找到最小值。将 dp[0] 初始化为 0，因为凑成金额 0 所需的硬币数为 0。

// 动态转移方程: 对于每个金额 i，我们尝试使用每个硬币 coin 来更新 dp[i] 的值。
// 如果 i >= coin，那么 dp[i] = min(dp[i], dp[i - coin] + 1)。

// 返回结果: 最后检查 dp[amount] 的值，如果它仍然是初始化时的那个大数值，
// 说明无法凑成该金额，返回 -1；否则返回 dp[amount]。

// 解释
// dp 数组的每个元素代表达到该金额所需的最小硬币数。
// 对每个金额 i，通过遍历所有硬币面额 coin，如果 i 可以由 coin 减去一些金额得到，我们尝试更新 dp[i] 的值。
// 最终返回 dp[amount] 的值。如果无法凑成，返回 -1。
// 时间复杂度
// 该算法的时间复杂度为 O(amount * n)，其中 n 是硬币的种类数。

func coinChange(coins []int, amount int) int {
	// 初始化 dp 数组
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	// 动态规划过程
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coin]+1)))
			}
		}
	}

	// 检查结果
	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}
