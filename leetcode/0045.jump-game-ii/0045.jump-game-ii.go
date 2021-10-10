package problem0045

// 45. 跳跃游戏 II
// 给你一个非负整数数组 nums ，你最初位于数组的第一个位置。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
// 假设你总是可以到达数组的最后一个位置。

// 方法一：贪心算法
func jump(nums []int) int {
	n := len(nums)
	// 当前覆盖范围，下一个覆盖范围
	curDistance, nextDistance, ans := 0, 0, 0
	for i := 0; i < n; i++ {
		// 下一个覆盖范围
		nextDistance = max(nums[i]+i, nextDistance)
		// 本覆盖范围的临界点
		if i == curDistance {
			if i == n-1 { //已经到了最后一个节点
				break
			}
			curDistance = nextDistance
			ans++
			// 已经覆盖了所有值
			if curDistance >= n-1 {
				break
			}
		}
	}
	return ans
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
