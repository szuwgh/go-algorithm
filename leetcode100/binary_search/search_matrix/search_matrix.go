package search_matrix

// SearchMatrix 在给定的矩阵中查找目标值

// 给你一个满足下述两条属性的 m x n 整数矩阵：

// 每行中的整数从左到右按非严格递增顺序排列。
// 每行的第一个整数大于前一行的最后一个整数。
// 给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。

// 示例 1：

// 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// 输出：true
// 示例 2：

// 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
// 输出：false
// 解释：
// 矩阵的虚拟一维数组：

// 我们把矩阵看作一个一维数组。对于矩阵中的每一个元素 matrix[row][col]，可以通过公式 index = row * cols + col 将它映射到一维数组中的索引。
// 二分查找：

// 使用标准的二分查找算法，通过不断调整 left 和 right 指针来查找目标值。
// 这种方法的时间复杂度是 O(log(m * n))，其中 m 是矩阵的行数，n 是列数。这样可以有效地利用矩阵的结构来提高查找效率。
func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	// 获取矩阵的行数和列数
	rows := len(matrix)
	cols := len(matrix[0])

	// 使用二分查找的变体
	left, right := 0, rows*cols-1

	for left <= right {
		// 计算中间位置
		mid := (left + right) / 2
		// 将一维索引转换为矩阵中的行列索引
		row := mid / cols
		col := mid % cols
		// 获取矩阵中对应位置的值
		midValue := matrix[row][col]

		if midValue == target {
			return true
		} else if midValue < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
