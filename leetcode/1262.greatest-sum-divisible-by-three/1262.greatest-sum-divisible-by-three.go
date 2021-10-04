package problem1262

// 1262. 可被三整除的最大和
// 给你一个整数数组 nums，请你找出并返回能被三整除的元素最大和。

func maxSumDivThree(nums []int) int {
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%3 == 0 {
		return sum
	}
	// dp[i][j] 表示前i个元素，除3余i的最大和
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, 3)
	}
	// INT_MAX = int((^uint(0)) >> 1)
	INT_MIN := ^int((^uint(0)) >> 1)
	dp[0][0] = 0
	dp[0][1] = INT_MIN
	dp[0][2] = INT_MIN
	var prei int
	var curi int
	for i := 1; i <= n; i++ {
		v := nums[i-1]
		// if v % 3 == 0 {
		//     dp[i][0] = max(dp[i-1][0], dp[i-1][0] + v)
		//     dp[i][1] = max(dp[i-1][1], dp[i-1][1] + v)
		//     dp[i][2] = max(dp[i-1][2], dp[i-1][2] + v)
		// } else if v % 3 == 1 {
		//     dp[i][0] = max(dp[i-1][0], dp[i-1][2] + v)
		//     dp[i][1] = max(dp[i-1][1], dp[i-1][0] + v)
		//     dp[i][2] = max(dp[i-1][2], dp[i-1][1] + v)
		// } else if v % 3 == 2 {
		//     dp[i][0] = max(dp[i-1][0], dp[i-1][1] + v)
		//     dp[i][1] = max(dp[i-1][1], dp[i-1][2] + v)
		//     dp[i][2] = max(dp[i-1][2], dp[i-1][0] + v)
		// }
		prei, curi = (i+1)%2, i%2
		for j := 0; j < 3; j++ {
			k := v % 3
			k = (j - k + 3) % 3
			dp[curi][j] = max(dp[prei][j], dp[prei][k]+v)
		}
	}
	return dp[curi][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
