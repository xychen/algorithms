package problem1035

// 1035. 不相交的线
// 在两条独立的水平线上按给定的顺序写下 nums1 和 nums2 中的整数。

// 现在，可以绘制一些连接两个数字 nums1[i] 和 nums2[j] 的直线，这些直线需要同时满足满足：

//  nums1[i] == nums2[j]
// 且绘制的直线不与任何其他连线（非水平线）相交。
// 请注意，连线即使在端点也不能相交：每个数字只能属于一条连线。

// 以这种方法绘制线条，并返回可以绘制的最大连线数。

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	if n1 == 0 || n2 == 0 {
		return 0
	}

	dp := make([][]int, n1+1)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n2+1)
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i%2][j] = dp[(i-1)%2][j-1] + 1
			} else {
				dp[i%2][j] = max(dp[(i-1)%2][j], dp[i%2][j-1])
			}
		}
	}
	return dp[n1%2][n2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
