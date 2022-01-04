package problem0343

// //回溯解法
// func integerBreak(n int) int {
//     var backtrack func(n int) int
//     backtrack = func(n int) int {
//         if n <= 2 {
//             return 1
//         }
//         m := 1
//         for i := 1; i < n; i++ {
//             m = max(m, i * backtrack(n-i), i * (n-i))
//         }
//         return m
//     }
//     return backtrack(n)
// }

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 1
	for i := 3; i <= n; i++ {
		for j := i - 1; j >= 1; j-- {
			dp[i] = max(dp[i], dp[j]*(i-j), (i-j)*j)
		}
	}
	return dp[n]
}

func max(nums ...int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}
