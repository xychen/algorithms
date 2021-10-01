package problem0718

// 718. 最长重复子数组
func findLength(nums1 []int, nums2 []int) int {
	n := len(nums1)
	m := len(nums2)
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, m+1)
		dp[i][0] = 0
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = 0
	}
	maxRes := 0
	for i := 1; i <= n; i++ {
		curi := i % 2
		prei := (i + 1) % 2
		for j := 1; j <= m; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[curi][j] = dp[prei][j-1] + 1
				maxRes = max(maxRes, dp[curi][j])
			} else {
				dp[curi][j] = 0
			}
		}
	}
	return maxRes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
