package problem0322

// 零钱兑换
// 方法一： 动态规划
func coinChange(coins []int, amount int) int {
	MAX_RES := amount + 1
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = MAX_RES
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin > i {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == MAX_RES {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
