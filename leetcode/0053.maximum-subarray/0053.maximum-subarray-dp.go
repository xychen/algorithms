package problem0053

// 最大子序和

// 方法二：动态规划 + 整个dp数组

func maxSubArray2(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	maxSum := 0
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxSum = dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		maxSum = max(maxSum, dp[i])
	}
	return maxSum
}
