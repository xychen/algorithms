package problem14

import "math"

// 贪心算法
func cuttingRope(n int) int {
	if n < 2 {
		return 0
	}
	if n <= 3 {
		return n - 1
	}

	a, b := n/3, n%3
	if b == 0 {
		return int(math.Pow(3, float64(a)))
	} else if b == 1 {
		return int(math.Pow(3, float64(a-1))) * 4
	} else if b == 2 {
		return int(math.Pow(3, float64(a))) * 2
	}
	return -1
}

// 方法二：动态规划版本
func cuttingRope2(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		// j是第1刀的位置
		for j := 2; j <= i; j++ {
			// j*(i-j) 表示只剪1刀
			// j*dp[i-j] 表示剪完1到，后边的长度（i-j）继续剪
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
