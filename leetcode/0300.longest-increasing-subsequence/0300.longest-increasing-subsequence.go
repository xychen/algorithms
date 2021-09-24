package problem0300

// 300. 最长递增子序列

//方法1： 动态规划
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	maxRes := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] <= nums[j] {
				continue
			}
			dp[i] = max(dp[i], dp[j]+1)
		}
		maxRes = max(maxRes, dp[i])
	}
	return maxRes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
