package problem0354

// . 俄罗斯套娃信封问题

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) <= 0 {
		return 0
	}
	//按w递增排序，w相同时再按h递减排序
	twoD := &TwoD{envelopes}
	sort.Sort(twoD)
	envelopes = twoD.data
	//按h求最长递增序列
	maxLen := 1
	dp := make([]int, len(envelopes))
	dp[0] = 1

	for i := 0; i < len(envelopes); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelopes[j][1] >= envelopes[i][1] {
				continue
			}
			dp[i] = max(dp[i], dp[j]+1)
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

type TwoD struct {
	data [][]int
}

func (twoD *TwoD) Len() int {
	return len(twoD.data)
}

func (twoD *TwoD) Less(i, j int) bool {
	if twoD.data[i][0] < twoD.data[j][0] {
		return true
	} else if twoD.data[i][0] > twoD.data[j][0] {
		return false
	}
	//data[i][0] == data[j][0]
	//h递减排序
	if twoD.data[i][1] > twoD.data[j][1] {
		return true
	} else {
		return false
	}
}

func (twoD *TwoD) Swap(i, j int) {
	twoD.data[i][0], twoD.data[j][0] = twoD.data[j][0], twoD.data[i][0]
	twoD.data[i][1], twoD.data[j][1] = twoD.data[j][1], twoD.data[i][1]
}
