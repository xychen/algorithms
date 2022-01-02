package problem0131

// 131. 分割回文串
// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

func partition(s string) [][]string {
	if len(s) <= 0 {
		return [][]string{}
	}
	res := make([][]string, 0)
	var backtrack func(s string, startIndex int, path []string)
	backtrack = func(s string, startIndex int, path []string) {
		if startIndex >= len(s) {
			tmpPath := make([]string, len(path))
			copy(tmpPath, path)
			res = append(res, tmpPath)
			return
		}

		for i := startIndex; i < len(s); i++ {
			if !isPlalindrome(s, startIndex, i) {
				continue
			}
			path = append(path, s[startIndex:i+1])
			backtrack(s, i+1, path)
			path = path[:len(path)-1]
		}
	}
	path := make([]string, 0)
	backtrack(s, 0, path)
	return res
}

// 判断回文方法可优化
func isPlalindrome(s string, left, right int) bool {
	for left <= right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
