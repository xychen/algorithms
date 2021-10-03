package problem1105

// 填充书架

func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	// dp[i]表示以第i本书结尾的书架的最小高度
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + books[i-1][1]
		tmpWidth, maxH := books[i-1][0], books[i-1][1]
		//从上往下拿
		for j := i - 2; j >= 0; j-- {
			tmpWidth += books[j][0]
			if tmpWidth > shelfWidth {
				break
			}
			maxH = max(maxH, books[j][1])
			dp[i] = min(dp[i], dp[j]+maxH)
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
