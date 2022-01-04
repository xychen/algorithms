package problem0416

// 分割等和子集
// 给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

// 方法一：dp+压缩空间
func canPartition(nums []int) bool {
	total, maxNum := 0, 0
	for _, num := range nums {
		total += num
		maxNum = max(maxNum, num)
	}
	// 和为奇数 或 存在某个元素大于 total/2
	if total%2 == 1 || maxNum > total/2 {
		return false
	}

	bagWeight := total / 2 //背包容量
	dp := make([]int, bagWeight+1)
	for j := nums[0]; j <= bagWeight; j++ {
		dp[j] = nums[0]
	}

	for i := 1; i < len(nums); i++ {
		for j := bagWeight; j >= 1; j-- {
			if j-nums[i] > 0 {
				dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
			}
		}
	}
	if dp[bagWeight] == bagWeight {
		return true
	}
	return false
}

// 方法二：动态规划 （未压缩空间）
func canPartition2(nums []int) bool {
	total, maxNum := 0, 0
	for _, num := range nums {
		total += num
		maxNum = max(maxNum, num)
	}
	// 和为奇数 或 存在某个元素大于 total/2
	if total%2 == 1 || maxNum > total/2 {
		return false
	}

	bagWeight := total / 2 //背包容量
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, bagWeight+1)
	}
	for j := nums[0]; j <= bagWeight; j++ {
		dp[0][j] = nums[0]
	}

	for i := 1; i < len(nums); i++ {
		for j := 1; j <= bagWeight; j++ {
			if j-nums[i] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]]+nums[i])
			}
		}
	}
	if dp[len(nums)-1][bagWeight] == bagWeight {
		return true
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 方法三：回溯 （会超时）
func canPartition3(nums []int) bool {
	total := sum(nums)
	if total%2 == 1 {
		return false
	}
	// 选择k个数，和等于 total/2
	var backtrack func(nums []int, startIndex, target int) bool
	backtrack = func(nums []int, startIndex, target int) bool {
		if target == 0 {
			return true
		}
		if startIndex >= len(nums) {
			return false
		}

		for i := startIndex; i < len(nums); i++ {
			if backtrack(nums, i+1, target-nums[i]) {
				return true
			}
		}
		return false
	}
	return backtrack(nums, 0, total/2)
}

func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
