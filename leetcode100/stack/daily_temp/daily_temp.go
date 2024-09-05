package daily_temp

// 每日温度
// 中等
// 提示
// 给定一个整数数组 temperatures ，表示每天的温度，
// 返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，
// 下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。

// 示例 1:

// 输入: temperatures = [73,74,75,71,69,72,76,73]
// 输出: [1,1,4,2,1,1,0,0]
// 示例 2:

// 输入: temperatures = [30,40,50,60]
// 输出: [1,1,1,0]
// 示例 3:

// 输入: temperatures = [30,60,90]
// 输出: [1,1,0]

// 提示：

// 1 <= temperatures.length <= 105
// 30 <= temperatures[i] <= 100

// 你可以通过使用 单调栈 来解决这个问题。对于每一天，
// 你可以检查未来哪一天的温度更高。我们维护一个存储下标的栈，栈中的元素对应的温度是从栈底到栈顶单调递减的。
// 这样在遍历温度数组时，每次遇到比栈顶温度高的温度时，就可以计算出当前温度是某些天之后的更高温度。

// 解释：
// 栈的作用：栈用来保存还未找到比当前温度高的那些天的索引。当我们找到更高温度时，弹出栈顶并计算这两天的距离。
// 时间复杂度：每个元素最多会被压入栈和弹出栈各一次，所以总的时间复杂度是 O(n)，其中 n 是温度数组的长度。
// 空间复杂度：由于使用了额外的栈，空间复杂度为 O(n)。
// 这段代码能在 O(n) 时间内解决问题，并且符合题目的约束。

// 具体步骤：
// 创建一个数组 answer 来保存结果，初始化为全 0。
// 创建一个栈 stack，用于存储尚未找到更高温度的天数索引。
// 从左到右遍历温度数组：
// 如果当前温度大于栈顶所存天数的温度，说明找到了一天比之前那一天温度更高的日子。
// 将栈顶元素弹出，并更新对应位置的 answer 值为两天之间的间隔天数。
// 重复上述过程，直到栈为空或栈顶温度不低于当前温度。
// 最后将当前天数的索引压入栈中。
// 遍历完成后，栈中的索引对应的天数后面没有比它们更高的温度，因此这些位置保持为 0。

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	answer := make([]int, n)
	stack := []int{} // 栈，用来存储还没找到更高温度的索引

	for i := 0; i < n; i++ {
		// 当当前温度大于栈顶所指的温度时，计算间隔天数并更新 answer
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[prevIndex] = i - prevIndex
		}
		// 将当前索引加入栈中
		stack = append(stack, i)
	}

	return answer
}
