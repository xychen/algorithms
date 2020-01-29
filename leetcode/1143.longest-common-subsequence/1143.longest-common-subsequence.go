package problem1143

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	if m == 0 || n == 0 {
		return 0
	}
	// m行， n 列， 默认初始值为0
	//dp := make([][n+1]int, m+1)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			a, b := i-1, j-1
			if text1[a] == text2[b] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

//优化数据表，只需要2行即可
func longestCommonSubsequence_2(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	if m == 0 || n == 0 {
		return 0
	}
	// m行， n 列， 默认初始值为0
	//只需要复用2行即可
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			a, b := i-1, j-1
			curRow := i % 2
			preRow := (i + 1) % 2
			if text1[a] == text2[b] {
				dp[curRow][j] = dp[preRow][j-1] + 1
			} else {
				dp[curRow][j] = max(dp[preRow][j], dp[curRow][j-1])
			}
		}
	}
	return dp[m%2][n]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
