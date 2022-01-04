package problem0494

// 给你一个整数数组 nums 和一个整数 target 。
// 向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
// 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
// 返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目

//解法思路： left - right = target  =>  left - (sum-left) = target  =>  left = (sum+target)/2

// 方法一：转成动态规划的0-1背包问题
func findTargetSumWays(nums []int, target int) int {
	total := sum(nums...)
	if total < target || (total+target)%2 == 1 {
		return 0
	}
	target = (total + target) / 2
	// 避免创建数组出现问题
	if target < 0 {
		return 0
	}
	// dp[j] 表示 填满j这么大容量的背包，有dp[j]种方法
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[target]
}

// 方法二： 回溯
func findTargetSumWays2(nums []int, target int) int {
	total := sum(nums...)
	if target > total || (total+target)%2 == 1 {
		return 0
	}
	target = (total + target) / 2

	res := 0
	var backtrack func(nums []int, startIndex, target int)
	backtrack = func(nums []int, startIndex, target int) {
		if target == 0 {
			res++
		}
		for i := startIndex; i < len(nums); i++ {
			// 做选择
			backtrack(nums, i+1, target-nums[i])
			// 回退
		}
	}
	backtrack(nums, 0, target)
	return res
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
