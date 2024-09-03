package longest

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长
// 子串
//  的长度。

// 示例 1:

// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:

// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:

// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//charSet 是一个哈希集合，用于存储当前窗口内的字符。
//如果遇到重复字符，我们会不断缩小窗口（移动左指针 left），直到将重复字符移出窗口。
//每次我们更新窗口后，检查当前窗口的长度是否是最大值。

func lengthOfLongestSubstring(s string) int {
	charSet := make(map[byte]bool)
	left, maxLength := 0, 0

	for right := 0; right < len(s); right++ {
		for charSet[s[right]] {
			delete(charSet, s[left])
			left++
		}
		charSet[s[right]] = true
		maxLength = max(maxLength, right-left+1)
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
