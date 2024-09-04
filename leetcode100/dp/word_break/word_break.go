package wordbreak

// 单词拆分
// 给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。

// 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。

// 示例 1：

// 输入: s = "leetcode", wordDict = ["leet", "code"]
// 输出: true
// 解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
// 示例 2：

// 输入: s = "applepenapple", wordDict = ["apple", "pen"]
// 输出: true
// 解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
//      注意，你可以重复使用字典中的单词。
// 示例 3：

// 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
// 输出: false

// 你可以使用动态规划（Dynamic Programming, DP）来解决这个问题。动态规划的核心思想是通过将问题分解成更小的子问题来求解整体问题。

// 具体思路如下：

// 创建一个 dp 数组，长度为 len(s) + 1。dp[i] 表示字符串 s 从起始位置到第 i 个字符是否可以由字典中的单词拼接而成。

// 初始条件为 dp[0] = true，表示空字符串总是可以由字典中的单词拼接而成。

// 遍历字符串 s 的每一个子字符串，并检查这个子字符串是否可以由字典中的单词拼接而成。如果可以，那么 dp[j] 也为 true。

// 如果最终 dp[len(s)] 为 true，则表示字符串 s 可以由字典中的单词拼接而成，否则不能。

// 解释：
// wordSet 是将 wordDict 转换为一个集合，用于快速查找单词是否存在。
// dp[i] 表示字符串 s 的前 i 个字符是否可以由字典中的单词拼接而成。
// 内层循环 j 遍历 s 的子字符串 s[j:i]，并检查这个子字符串是否在 wordSet 中。
// 复杂度分析：
// 时间复杂度：O(n^2)，其中 n 是字符串 s 的长度。外层循环运行 n 次，内层循环最多也运行 n 次。
// 空间复杂度：O(n)，dp 数组的长度为 n+1。
// 这个算法能够有效地解决此类问题，适用于大部分常见情况。

func wordBreak(s string, wordDict []string) bool {
	// 将 wordDict 转换为一个集合，以便快速查找
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	// dp 数组，dp[i] 表示 s[0:i] 是否可以由 wordDict 中的单词拼接成
	dp := make([]bool, len(s)+1)
	dp[0] = true // 空字符串可以由字典中的单词拼接而成

	// 遍历 s 的每一个位置
	for i := 1; i <= len(s); i++ {
		// 再从位置 j 切分字符串
		for j := 0; j < i; j++ {
			// 如果 dp[j] 是 true，并且 s[j:i] 在字典中
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}
