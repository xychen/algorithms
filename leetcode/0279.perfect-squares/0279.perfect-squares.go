package problem0279

// 完全平方数

import "math"

func numSquares(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		m := int(math.Floor(math.Sqrt(float64(i))))
		dp[i] = i
		for j := 1; j <= m; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
