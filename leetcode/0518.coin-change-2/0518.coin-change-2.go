package problem0518

func change(amount int, coins []int) int {
	// dp[j]的含义：兑换金额为 amount 的时候，有多少种组合，等于所有 dp[j-coin] 之和
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] += dp[j-coin]
		}
	}
	return dp[amount]
}
