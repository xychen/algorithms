package problem0188

// 188. 买卖股票的最佳时机 IV
// 给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。

// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。

// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

func maxProfit(k int, prices []int) int {
	if k < 1 || len(prices) <= 0 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2*k+1)
	}
	// 初始化
	for j := 0; j < 2*k+1; j++ {
		if j%2 == 1 {
			dp[0][j] = -prices[0]
		}
	}

	// dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	// dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
	// dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
	// dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])

	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		for j := 1; j < 2*k+1; j++ {
			if j%2 == 1 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]+prices[i])
			}
		}
	}
	return dp[len(prices)-1][2*k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
