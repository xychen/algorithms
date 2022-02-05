package problem0045

// 45. 跳跃游戏 II
// 给你一个非负整数数组 nums ，你最初位于数组的第一个位置。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
// 假设你总是可以到达数组的最后一个位置。

// 方法一：贪心算法
func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	res, cover, nextCover := 1, nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if cover >= len(nums)-1 {
			break
		}
		nextCover = max(nextCover, i+nums[i])
		if i >= cover {
			res++
			cover = nextCover
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

// 方法二：动态规划
func jump2(nums []int) int {
	// dp[i] 表示到i下标使用的最少跳数
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		dp[i] = n + 1
		for j := i - 1; j >= 0; j-- {
			if nums[j]+j >= i {
				dp[i] = min(dp[j]+1, dp[i])
			}
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
