package problem0377

// 377. 组合总和 Ⅳ
// 给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。

// 方法一：动态规划 （零钱兑换问题）
// 本题求排列数
// 如果求组合数，外层for循环遍历物品、内层for循环遍历背包
// 如果求排列数，外层for循环遍历背包、内层for循环遍历物品
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
