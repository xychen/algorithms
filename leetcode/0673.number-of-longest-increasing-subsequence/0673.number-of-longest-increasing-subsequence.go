package problem0673

// 最长递增子序列的个数

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	// dp[i] 表示以nums[i] 结束的递增子序列的最大值
	dp := make([]int, n)
	// dpCnt[i] 表示以nums[i] 结尾的递增子序列的最大值的个数
	dpCnt := make([]int, n)
	dp[0] = 1
	dpCnt[0] = 1
	maxLen := 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		dpCnt[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] <= nums[j] {
				continue
			}
			//存在递增情况
			if dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				dpCnt[i] = dpCnt[j]
			} else if dp[j]+1 == dp[i] {
				dpCnt[i] += dpCnt[j]
			}

		}
		maxLen = max(dp[i], maxLen)
	}
	// fmt.Println(maxLen, ",", dp, ",", dpCnt)
	res := 0
	for i := 0; i < n; i++ {
		if maxLen == dp[i] {
			res += dpCnt[i]
		}
	}
	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
