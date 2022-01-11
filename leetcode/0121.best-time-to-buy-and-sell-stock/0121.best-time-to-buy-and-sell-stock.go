package problem0121

// 121. 买卖股票的最佳时机
// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

func maxProfit(prices []int) int {
	// dp[i][0] 表示第i天持有股票的最大收益， dp[i][1]表示不持有股票的最大收益
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		// 第i天持有股票：max(第i天购买股票，前一天就购买了股票)
		dp[i][0] = max(-prices[i], dp[i-1][0])
		// 第i天不持有股票：max(第i天卖出了股票，前一天就卖出了股票)
		dp[i][1] = max(dp[i-1][0]+prices[i], dp[i-1][1])
	}
	// 最后一天肯定是不持有股票是最大收益
	return dp[len(prices)-1][1]
}

// 优化空间
func maxProfit2(prices []int) int {
	// dp[i][0] 表示第i天持有股票的最大收益， dp[i][1]表示不持有股票的最大收益
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		// 第i天持有股票：max(第i天购买股票，前一天就购买了股票)
		dp[i%2][0] = max(-prices[i], dp[(i-1)%2][0])
		// 第i天不持有股票：max(第i天卖出了股票，前一天就卖出了股票)
		dp[i%2][1] = max(dp[(i-1)%2][0]+prices[i], dp[(i-1)%2][1])
	}
	// 最后一天肯定是不持有股票是最大收益
	return dp[(len(prices)-1)%2][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
