package problem0053

// 最大子序和

// 方法三：递归+记忆化搜索

var maxSum int
var dp []int

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func maxSubArray3(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	maxSum = 0
	dp = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = INT_MIN
	}
	dp[0] = nums[0]
	maxSum = dp[0]
	for i := 1; i < len(nums); i++ {
		tmpRes := solve(nums, i)
		maxSum = max(maxSum, tmpRes)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(nums []int, i int) int {
	if i == 0 {
		return nums[i]
	}
	if dp[i] != INT_MIN {
		return dp[i]
	}
	dp[i] = max(solve(nums, i-1)+nums[i], nums[i])
	return dp[i]
}
