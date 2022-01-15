package problem0115

// 115. 不同的子序列
// 给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。
// 字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE" 的一个子序列，而 "AEC" 不是）
// 题目数据保证答案符合 32 位带符号整数范围。

func numDistinct(s string, t string) int {
    n1, n2 := len(s), len(t)
    dp := make([][]int, n1+1)
    for i := 0; i <= n1; i++ {
        dp[i] = make([]int, n2+1)
        dp[i][0] = 1
    }

    for i := 1; i <= n1; i++ {
        for j := 1; j <= n2; j++ {
            if s[i-1] == t[j-1] {
                dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
            } else {
                dp[i][j] = dp[i-1][j]
            }
        }
    }
    return dp[n1][n2]
}