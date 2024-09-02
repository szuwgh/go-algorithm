package longest_palindromic_substring

// 5.最长回文子串
// 题目描述：
// 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

// 示例 1：
// 输入: "babad"
// 输出: "bab"
// 注意: "aba" 也是一个有效答案。
// 1
// 2
// 3

// 示例 2：
// 输入: "cbbd"
// 输出: "bb"

func longestPalinddrome(s string) string {
	if len(s) < 2 {
		return s
	}
	//最长回文的首字符索引，最长回文的长度
	begin, maxLen := 0, 1
	for i := 0; i < len(s); {
		if len(s)-i <= maxLen/2 {
			break
		}
		// b 代表回文的**首**字符索引号，
		// e 代表回文的**尾**字符索引号，
		b, e := i, i
		for e < len(s)-1 && s[e+1] == s[e] {
			e++
			// 循环结束后，s[b:e+1]是一串相同的字符串
		}
		i = e + 1
		// 先从i开始，利用`正中间段`所有字符相同的特性，让b，e分别指向`正中间段`的首尾
		// 再从`正中间段`向两边扩张，让b，e分别指向此`正中间段`为中心的最长回文的首尾
		for e < len(s)-1 && b > 0 && s[e+1] == s[b-1] {
			e++
			b--
			// 循环结束后，s[b:e+1]是这次能找到的最长回文。
		}
		newLen := e + 1 - b
		if newLen > maxLen {
			begin = b
			maxLen = newLen
		}
	}
	return s[begin : begin+maxLen]
}
