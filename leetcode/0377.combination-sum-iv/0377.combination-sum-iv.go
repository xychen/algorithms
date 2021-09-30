package problem0377

// 377. 组合总和 Ⅳ
// 给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。

// 方法一：动态规划 （零钱兑换问题）
func combinationSum4(nums []int, target int) int {
	// dp[i] 表示和为i的元素组合个数
	dp := make([]int, target+1)
	//不取任何元素时，和才能为0，所以dp[0]=1
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

// 方法二： 回溯 + 记忆化搜索
var dp map[int]int

func combinationSum4_2(nums []int, target int) int {
	dp = make(map[int]int)
	return solve(nums, target)
}

func solve(nums []int, target int) int {
	if target == 0 {
		return 1
	}
	if target < 0 {
		return 0
	}
	if v, ok := dp[target]; ok {
		return v
	}
	res := 0
	for _, num := range nums {
		//最选择
		res += solve(nums, target-num)
		//回退
	}
	dp[target] = res
	return res
}
