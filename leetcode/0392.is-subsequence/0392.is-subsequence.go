package problem0392

// 392. 判断子序列
// 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
// 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
// 进阶：
// 如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？

func isSubsequence(s string, t string) bool {
	n1, n2 := len(s), len(t)
	if n1 == 0 {
		return true
	}
	if n2 == 0 {
		return false
	}

	dp := make([][]int, n1+1)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n2+1)
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if s[i-1] == t[j-1] {
				dp[i%2][j] = dp[(i-1)%2][j-1] + 1
			} else {
				dp[i%2][j] = max(dp[i%2][j-1], dp[(i-1)%2][j])
			}
		}
	}
	return dp[n1%2][n2] == n1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
