package max_profit

// 121. 买卖股票的最佳时机
// 简单
// 相关标签
// 相关企业
// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

// 示例 1：

// 输入：[7,1,5,3,6,4]
// 输出：5
// 解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//      注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
// 示例 2：

// 输入：prices = [7,6,4,3,1]
// 输出：0
// 解释：在这种情况下, 没有交易完成, 所以最大利润为 0。

// 要解决这个问题，我们可以使用一个简单的线性扫描算法。这个算法的核心思想是：

// 维护一个变量 minPrice 来记录到当前为止的最小股票价格。
// 维护一个变量 maxProfit 来记录当前的最大利润。
// 遍历股票价格数组，对每一天的价格进行比较：
// 如果当前价格小于 minPrice，更新 minPrice。
// 计算当前价格与 minPrice 的差值，如果这个差值大于 maxProfit，更新 maxProfit。
// 这样，我们可以在一次遍历中找到最大利润。这个算法的时间复杂度是 O(n)，空间复杂度是 O(1)。

// 代码说明
// maxProfit 函数接受一个整数数组 prices 作为参数，表示股票的价格。
// minPrice 初始为数组的第一个元素，它代表了到目前为止的最低价格。
// maxProfit 用于记录最大利润，初始化为 0。
// 遍历 prices 数组，如果当前价格小于 minPrice，更新 minPrice。计算当前价格与 minPrice 的差值，并更新 maxProfit 如果当前利润更大。
// main 函数中测试了两个示例，输出对应的最大利润。
// 这种方法有效且高效，适用于任何股票价格数组。

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}
		profit := price - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}
