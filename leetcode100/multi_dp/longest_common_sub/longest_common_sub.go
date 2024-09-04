package long

// 最长公共子序列
// 给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。

// 一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。

// 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
// 两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。

// 示例 1：

// 输入：text1 = "abcde", text2 = "ace"
// 输出：3
// 解释：最长公共子序列是 "ace" ，它的长度为 3 。
// 示例 2：

// 输入：text1 = "abc", text2 = "abc"
// 输出：3
// 解释：最长公共子序列是 "abc" ，它的长度为 3 。
// 示例 3：

// 输入：text1 = "abc", text2 = "def"
// 输出：0
// 解释：两个字符串没有公共子序列，返回 0 。

// 要在 Go 中解决这个问题，可以使用动态规划（Dynamic Programming，DP）。
// 我们将创建一个二维的 DP 表格，表格的大小为 len(text1) + 1 行乘以 len(text2) + 1 列。
// dp[i][j] 表示 text1[0:i] 和 text2[0:j] 的最长公共子序列的长度。

// 动态规划的思路：
// 如果 text1[i-1] == text2[j-1]，则 dp[i][j] = dp[i-1][j-1] + 1。
// 这是因为如果两个字符相等，那么最长公共子序列的长度在之前的基础上加1。

// 如果 text1[i-1] != text2[j-1]，则 dp[i][j] = max(dp[i-1][j], dp[i][j-1])。
// 这是因为我们需要考虑删除 text1 的一个字符或者 text2 的一个字符，看哪个删除后可以得到更长的公共子序列。

// 初始化的时候，dp[i][0] 和 dp[0][j] 都应该是0，因为如果任意一个字符串是空的，那么最长公共子序列的长度为0。

// 最终结果保存在 dp[len(text1)][len(text2)] 中。

// 解释：
// dp[i][j] 是 text1 的前 i 个字符和 text2 的前 j 个字符的最长公共子序列长度。
// 当字符匹配时，当前长度比上一个子问题多 1。
// 当字符不匹配时，取前一个状态的最大值。
// 这个解决方案的时间复杂度是
// 𝑂
// (
// 𝑚
// ×
// 𝑛
// )
// O(m×n)，空间复杂度也是
// 𝑂
// (
// 𝑚
// ×
// 𝑛
// )
// O(m×n)，其中 m 和 n 分别是 text1 和 text2 的长度。

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)

	// 创建 DP 数组
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 填充 DP 数组
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

// 辅助函数，返回两个整数中的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
