package maxProduct

import "math"

// 乘积最大子数组

// 给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续
// 子数组
// （该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

// 测试用例的答案是一个 32-位 整数。

// 示例 1:

// 输入: nums = [2,3,-2,4]
// 输出: 6
// 解释: 子数组 [2,3] 有最大乘积 6。
// 示例 2:

// 输入: nums = [-2,0,-1]
// 输出: 0
// 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。

// 代码解释：
// 初始状态：我们将 curMax 和 curMin 设为数组的第一个元素，代表以第一个元素结尾的最大乘积和最小乘积。
// 遍历数组：从第二个元素开始遍历：
// 如果当前数字是负数，会交换 curMax 和 curMin，因为负数乘以最小的数可能变成最大的数。
// 更新 curMax 为当前元素本身或者当前元素乘以之前的 curMax，取两者的最大值。
// 更新 curMin 为当前元素本身或者当前元素乘以之前的 curMin，取两者的最小值。
// 更新全局最大值：每次计算出新的 curMax 后，与 globalMax 进行比较，更新全局最大乘积。
// 返回结果：最终的 globalMax 即为数组中乘积最大的连续子数组的乘积。
// 示例解释：
// 对于输入 [2, 3, -2, 4]，最大乘积是 6，来自子数组 [2, 3]。
// 对于输入 [-2, 0, -1]，最大乘积是 0，因为任何子数组乘以 0 后结果都是 0。

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 初始化最大值、最小值以及全局最大值
	curMax, curMin, globalMax := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			// 如果当前元素是负数，交换当前最大值和最小值
			curMax, curMin = curMin, curMax
		}

		// 更新当前最大值和最小值
		curMax = int(math.Max(float64(num), float64(curMax*num)))
		curMin = int(math.Min(float64(num), float64(curMin*num)))

		// 更新全局最大值
		globalMax = int(math.Max(float64(globalMax), float64(curMax)))
	}

	return globalMax
}
