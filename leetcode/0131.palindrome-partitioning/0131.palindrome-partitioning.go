package problem0131

// 131. 分割回文串
// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

import "strings"

func partition(s string) [][]string {
	res := make([][]string, 0)
	path := make([]string, 0)
	mem := make([][]int, len(s))
	for i := range mem {
		mem[i] = make([]int, len(s))
	}

	var isPalindrome func(i, j int) int
	isPalindrome = func(i, j int) int {
		if i >= j {
			return 1
		}
		if mem[i][j] != 0 {
			return mem[i][j]
		}
		mem[i][j] = -1
		if s[i] == s[j] {
			mem[i][j] = isPalindrome(i+1, j-1)
		}
		return mem[i][j]
	}

	var solve func(start int)
	solve = func(start int) {
		str := strings.Join(path, "")
		if len(str) == len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := start; i < len(s); i++ {
			// if !ishuiwen(s, start, i) {
			//     continue
			// }

			// 使用记忆化搜索
			if isPalindrome(start, i) == -1 {
				continue
			}
			path = append(path, s[start:i+1])
			solve(i + 1)
			path = path[:len(path)-1]
		}
	}
	solve(0)
	return res
}

// 最原始的方法
func ishuiwen(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
