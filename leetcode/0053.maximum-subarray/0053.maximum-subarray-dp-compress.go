package problem0053

// 最大子序和

// 方法一：递归+记忆化搜索

func maxSubArray(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	pre := nums[0]
	cur := nums[0]
	maxSum := pre
	for i := 1; i < len(nums); i++ {
		cur = max(pre+nums[i], nums[i])
		maxSum = max(maxSum, cur)
		pre = cur
	}
	return maxSum
}
