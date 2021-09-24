package problem0300

// 300. 最长递增子序列

//方法二： 递归 + 记忆化搜索
var dp []int

func lengthOfLIS2(nums []int) int {
	dp = make([]int, len(nums))
	dp[0] = 1
	maxRes := 1
	for i := 0; i < len(nums); i++ {
		maxRes = max(solve(nums, i), maxRes)
	}
	return maxRes
}

func solve(nums []int, i int) int {
	if dp[i] > 0 {
		return dp[i]
	}
	res := 1
	for j := 0; j < i; j++ {
		if nums[i] <= nums[j] {
			continue
		}
		res = max(res, solve(nums, j)+1)
	}
	dp[i] = res
	return res
}
