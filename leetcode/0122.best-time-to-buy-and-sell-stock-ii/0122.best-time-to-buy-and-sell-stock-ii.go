package problem0122

// 122. 买卖股票的最佳时机 II
// 给定一个数组 prices ，其中 prices[i] 是一支给定股票第 i 天的价格。

// 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

// 分步累加利润
// https://programmercarl.com/0122.%E4%B9%B0%E5%8D%96%E8%82%A1%E7%A5%A8%E7%9A%84%E6%9C%80%E4%BD%B3%E6%97%B6%E6%9C%BAII.html
func maxProfit(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		ans += max(prices[i]-prices[i-1], 0)
	}
	return ans
}

// 动态规划解法
func maxProfit2(prices []int) int {
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
		dp[i%2][1] = max(dp[(i-1)%2][0]+prices[i], dp[(i-1)%2][1])
	}
	return dp[(len(prices)-1)%2][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
