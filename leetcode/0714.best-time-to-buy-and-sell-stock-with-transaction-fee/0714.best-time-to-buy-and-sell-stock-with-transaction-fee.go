package problem0714

// 714. 买卖股票的最佳时机含手续费
// 给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；整数 fee 代表了交易股票的手续费用。

// 你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。

// 返回获得利润的最大值。

// 注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。

func maxProfit(prices []int, fee int) int {
	// dp[i][0]表示第i天持有股票的最大收益，dp[i][1]表示第i天不持有股票的最大收益
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		// 持有股票：当天买的(加上原来的收益)，或者前一天就持有
		dp[i%2][0] = max(-prices[i]+dp[(i-1)%2][1], dp[(i-1)%2][0])
		// 未持有股票：当天卖了，或者前一天就没持有
		dp[i%2][1] = max(dp[(i-1)%2][0]+prices[i]-fee, dp[(i-1)%2][1])
	}
	return dp[(len(prices)-1)%2][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
