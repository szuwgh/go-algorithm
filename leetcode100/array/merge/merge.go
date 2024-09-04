package merge

import "sort"

// 合并区间
// 中等
// 相关标签
// 相关企业
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

// 示例 1：

// 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
// 输出：[[1,6],[8,10],[15,18]]
// 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
// 示例 2：

// 输入：intervals = [[1,4],[4,5]]
// 输出：[[1,5]]
// 解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

// 代码解析：
// 排序：sort.Slice 函数根据区间的起始位置对所有区间进行排序。
// 合并逻辑：
// 初始化合并结果 merged，将第一个区间放入。
// 遍历排序后的区间，如果当前区间的起始位置小于或等于上一个已合并区间的结束位置，说明有重叠，更新结束位置为二者之间的最大值。
// 如果没有重叠，则将当前区间直接添加到合并结果中。
// 输出：最终返回合并后的区间数组。
// 该实现的时间复杂度主要由排序决定，为 O(n log n)，其中 n 是区间的数量。

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	// 先按照每个区间的起始位置进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 用于存储合并后的区间结果
	merged := [][]int{}

	// 初始化第一个区间为待合并的区间
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		// 获取最后一个已经合并的区间
		lastMerged := merged[len(merged)-1]

		// 如果当前区间与最后一个合并的区间有重叠，更新合并后的区间的结束位置
		if intervals[i][0] <= lastMerged[1] {
			if intervals[i][1] > lastMerged[1] {
				lastMerged[1] = intervals[i][1]
			}
		} else {
			// 如果不重叠，直接将当前区间加入到结果中
			merged = append(merged, intervals[i])
		}
	}

	return merged
}
