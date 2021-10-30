package problem0055

// 方法一：贪心算法，好理解的版本
func canJump(nums []int) bool {
	// 覆盖范围
	cover := 0
	// 注意：终止条件是 i<=cover
	for i := 0; i <= cover; i++ {
		cover = max(cover, nums[i]+i)
		// 如果覆盖范围超过了最后一个元素，返回true
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}

// 方法一：贪心算法
func canJump2(nums []int) bool {
	n := len(nums)
	farthest := 0
	for i := 0; i < n-1; i++ {
		farthest = max(farthest, nums[i]+i)
		// 到i之后，加上nums[i]，理论上要超过i，未超过说明卡主了
		if farthest <= i {
			return false
		}
	}
	if farthest >= n-1 {
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

// 方法二：动态规划
func canJump3(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n)
	dp[0] = true
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && j+nums[j] >= i {
				dp[i] = true
				break
			}
		}
	}
	return dp[n-1]
}
