package problem0070

// 爬楼梯

func climbStairs(n int) int {
	// dp := make([]int, n+1)
	// dp[1] = 1
	// dp[2] = 2
	if n <= 2 {
		return n
	}
	pre, cur, tmp := 1, 2, 0
	for i := 3; i <= n; i++ {
		// dp[i] = dp[i-1] + dp[i-2]
		tmp = pre + cur
		pre, cur = cur, tmp
	}
	return cur
}
