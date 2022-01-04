package problem0096

// 不同的二叉搜索树

// 回溯解法(超时)
// func numTrees(n int) int {
//     var backtrack func(n int) int
//     backtrack = func(n int) int {
//         if n <= 1 {
//             return 1
//         }
//         res := 0
//         // 以i作为根节点
//         for i := 1; i <= n; i++ {
//             res += backtrack(i-1) * backtrack(n-i)
//         }
//         return res
//     }
//     return backtrack(n)
// }

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			//包含左子树为空或者右子树为空的情况
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
