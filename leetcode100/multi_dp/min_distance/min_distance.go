package min

// 编辑距离
// 给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

// 你可以对一个单词进行如下三种操作：

// 插入一个字符
// 删除一个字符
// 替换一个字符

// 示例 1：

// 输入：word1 = "horse", word2 = "ros"
// 输出：3
// 解释：
// horse -> rorse (将 'h' 替换为 'r')
// rorse -> rose (删除 'r')
// rose -> ros (删除 'e')
// 示例 2：

// 输入：word1 = "intention", word2 = "execution"
// 输出：5
// 解释：
// intention -> inention (删除 't')
// inention -> enention (将 'i' 替换为 'e')
// enention -> exention (将 'n' 替换为 'x')
// exention -> exection (将 'n' 替换为 'c')
// exection -> execution (插入 'u')

// 解释：
// 初始化 dp 数组：

// dp[i][0] 表示将 word1 的前 i 个字符转换为空字符串的操作数，也就是 i（全删除操作）。
// dp[0][j] 表示将空字符串转换成 word2 的前 j 个字符的操作数，也就是 j（全插入操作）。
// 填充 dp 数组：

// 如果 word1[i-1] 和 word2[j-1] 相等，那么 dp[i][j] 就是 dp[i-1][j-1]。
// 如果不相等，我们需要考虑三种操作：插入、删除、替换，取这三者的最小值加一。
// 返回结果：

// dp[m][n] 就是将 word1 转换为 word2 的最少操作数。

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	// 创建二维数组 dp
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 初始化 dp 数组
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// 填充 dp 数组
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}

	return dp[m][n]
}
