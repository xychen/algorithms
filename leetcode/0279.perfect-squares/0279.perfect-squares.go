package problem0279

// 完全平方数

func numSquares(n int) int {
	coins := make([]int, 0)
	for i := 1; i*i <= n; i++ {
		coins = append(coins, i*i)
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = n + 1
	}
	for _, coin := range coins {
		for i := coin; i <= n; i++ {
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	return dp[n]
}

// func numSquares(n int) int {
//	   //可以不用提前生成coin
//     dp := make([]int, n+1)
//     for i := 1; i <= n; i++ {
//         dp[i] = n+1
//     }
//     for i := 1; i*i <= n; i++ {
//         for j := i*i; j <= n; j++ {
//             dp[j] = min(dp[j], dp[j-i*i]+1)
//         }
//     }
//     return dp[n]
// }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
