package problem0583

// 583. 两个字符串的删除操作
// 给定两个单词 word1 和 word2，找到使得 word1 和 word2 相同所需的最小步数，每步可以删除任意一个字符串中的一个字符。

func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
		dp[i][0] = i
	}
	for j := 0; j <= n2; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1])+1, dp[i-1][j-1]+2)
			}
		}
	}
	return dp[n1][n2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
