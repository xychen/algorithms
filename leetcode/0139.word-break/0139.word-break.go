package problem0139

import "strings"

//动态规划解法
func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	if len(wordDict) == 0 {
		return false
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		dp[i] = false
		for _, word := range wordDict {
			//从i往前推len(world)
			if i-len(word) >= 0 && s[i-len(word):i] == word && dp[i-len(word)] {
				dp[i] = true
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
