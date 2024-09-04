package long

// 给你一个字符串 s，找到 s 中最长的回文子串

// 示例 1：

// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
// 示例 2：

// 输入：s = "cbbd"
// 输出："bb"

// 在 Go 语言中，你可以使用动态规划或中心扩展的方法来找到字符串 s 中的最长回文子串。
// 下面我将展示一个使用中心扩展方法的解决方案。

// 中心扩展的方法的核心思想是，每个字符和每对相邻字符都可能是回文的中心。
// 通过从每个中心向两边扩展，可以找到以该中心为中心的最长回文子串。

// 代码解析：
// longestPalindrome 函数：

// 遍历字符串的每个字符，并将其作为中心来扩展寻找回文。
// len1 代表以单个字符为中心的回文长度，len2 代表以两个字符之间的空隙为中心的回文长度。
// 根据计算出的最大回文长度来更新回文的起始和结束位置。
// expandAroundCenter 函数：

// 从给定的中心位置（left 和 right）开始向外扩展，直到遇到不同的字符或者超出字符串的边界。
// 返回扩展后的回文长度。
// max 函数：

// 用于求两个数中的最大值。
// 运行结果：
// 对于输入 "babad"，输出可能是 "bab" 或 "aba"，因为这两个回文子串的长度相同。
// 对于输入 "cbbd"，输出为 "bb"，因为这是最长的回文子串。

// 寻找最长回文子串
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		maxLen := max(len1, len2)
		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}
	return s[start : end+1]
}

// 辅助函数：从中心向两边扩展寻找回文
func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

// 辅助函数：求最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
