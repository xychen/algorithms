package problem0309

// 309. 最佳买卖股票时机含冷冻期
// 给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​

// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:

// 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。

// 题解：http://programmercarl.com/0309.最佳买卖股票时机含冷冻期.html#思路

func maxProfit(prices []int) int {
	// 状态一：买入股票状态（今天买入股票，或者是之前就买入了股票然后没有操作）
	// 卖出股票状态，这里就有两种卖出股票状态
	// 状态二：两天前就卖出了股票，度过了冷冻期，一直没操作，今天保持卖出股票状态
	// 状态三：今天卖出了股票
	// 状态四：今天为冷冻期状态，但冷冻期状态不可持续，只有一天！
	n := len(prices)
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 4)
	}
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][1], dp[i-1][3])-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}
	return max(dp[n-1][1], max(dp[n-1][2], dp[n-1][3]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
