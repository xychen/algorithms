package problem0139

// 单词拆分

import "strings"

//动态规划解法
// 完全背包问题求排列，所以背包重量在外层循环
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for _, word := range wordDict {
			if i < len(word) {
				continue
			}
			if s[i-len(word):i] == word && dp[i-len(word)] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

//回溯 + 记忆化搜索
var dp map[string]bool

func wordBreak2(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	if len(wordDict) == 0 {
		return false
	}
	dp = make(map[string]bool)
	return solve(s, wordDict)
}

func solve(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	if v, ok := dp[s]; ok {
		return v
	}
	res := false
	for _, word := range wordDict {
		if strings.Index(s, word) == 0 {
			//最选择
			res = solve(s[len(word):], wordDict)
			if res {
				break
			}
			//回退
		}
	}
	dp[s] = res
	return res
}
