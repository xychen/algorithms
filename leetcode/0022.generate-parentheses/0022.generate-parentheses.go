package problem0022

import "strings"

// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
// 有效括号组合需满足：左括号必须以正确的顺序闭合。

var res []string

func generateParenthesis(n int) []string {
	res = make([]string, 0)
	path := make([]string, 0)
	solve(n, n, path)
	return res
}

func solve(left, right int, path []string) {
	if left < 0 || right < 0 {
		return
	}
	if left == 0 && right == 0 {
		str := strings.Join(path, "")
		res = append(res, str)
		return
	}
	//不合法
	if left > right {
		return
	}

	//做选择
	path = append(path, "(")
	solve(left-1, right, path)
	//回退
	path = path[:len(path)-1]

	//做选择
	path = append(path, ")")
	solve(left, right-1, path)
	// 回退
	path = path[:len(path)-1]
}
