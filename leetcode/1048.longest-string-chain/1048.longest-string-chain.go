package problem1048

// 1048. 最长字符串链

import "sort"

func longestStrChain(words []string) int {
	n := len(words)
	if n <= 1 {
		return n
	}
	sort.Sort(MyStringList(words))
	//dp[i]表示words[i]为结尾单词的词链长度
	dp := make([]int, n)
	dp[0] = 1
	maxLen := 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if isAE(words[j], words[i]) {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLen = max(maxLen, dp[i])
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isAE(a, b string) bool {
	if len(b)-len(a) != 1 {
		return false
	}
	diff := 0
	i, j := 0, 0
	for i < len(a) {
		if a[i] != b[j] {
			diff++
			if diff >= 2 {
				return false
			}
			j++
			continue
		} else {
			i++
			j++
		}
	}
	return true
}

type MyStringList []string

func (p MyStringList) Len() int {
	return len(p)
}
func (p MyStringList) Less(i, j int) bool {
	return len(p[i]) < len(p[j])
}
func (p MyStringList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
