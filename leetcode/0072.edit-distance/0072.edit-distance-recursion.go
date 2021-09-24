package problem0072

// 编辑距离
// 方法一：递归+记忆搜索

var dp [][]int

const INT_MAX = int(^uint(0) >> 1)

func minDistance2(word1 string, word2 string) int {
	dp = make([][]int, len(word1))
	for i := 0; i < len(word1); i++ {
		dp[i] = make([]int, len(word2))
		for j := 0; j < len(word2); j++ {
			dp[i][j] = INT_MAX
		}
	}
	return solve(word1, word2, len(word1)-1, len(word2)-1)
}

func solve(word1, word2 string, i, j int) int {
	if i < 0 {
		return j + 1
	}
	if j < 0 {
		return i + 1
	}
	if dp[i][j] != INT_MAX {
		return dp[i][j]
	}
	res := INT_MAX
	if word1[i] == word2[j] {
		res = solve(word1, word2, i-1, j-1)
	} else {
		//插入、删除、替换
		res = min(solve(word1, word2, i, j-1)+1, solve(word1, word2, i-1, j)+1, solve(word1, word2, i-1, j-1)+1)
	}
	dp[i][j] = res
	return res
}

func min(nums ...int) int {
	minNum := nums[0]
	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
	}
	return minNum
}
