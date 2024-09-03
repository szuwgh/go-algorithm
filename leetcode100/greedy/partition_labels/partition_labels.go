package partition_labels

// 763. 划分字母区间
// 中等
// 相关标签
// 相关企业
// 提示
// 给你一个字符串 s 。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。

// 注意，划分结果需要满足：将所有划分结果按顺序连接，得到的字符串仍然是 s 。

// 返回一个表示每个字符串片段的长度的列表。

// 代码解释
// 存储每个字符的最后出现位置：

// 遍历字符串，使用一个映射 lastOccurrence 来记录每个字符的最后出现位置。
// 贪心划分片段：

// 初始化 start 和 end 变量来表示当前片段的开始和结束位置。
// 遍历字符串中的每个字符，更新 end 为当前字符的最后出现位置。
// 当遍历到字符的索引等于 end 时，说明当前片段划分完成，将其长度加入结果列表，并更新 start 为下一个片段的开始位置。
// 返回结果：

// 返回存储片段长度的列表。

func partitionLabels(s string) []int {
	// 存储每个字符最后出现的位置
	lastOccurrence := make(map[rune]int)
	for i, ch := range s {
		lastOccurrence[ch] = i
	}

	result := []int{}
	start, end := 0, 0

	for i, ch := range s {
		// 更新当前片段的结束位置
		if lastOccurrence[ch] > end {
			end = lastOccurrence[ch]
		}

		// 当当前字符的索引等于片段的结束位置时，划分完成
		if i == end {
			result = append(result, end-start+1)
			start = i + 1
		}
	}

	return result
}
