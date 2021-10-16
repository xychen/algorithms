package problem0216

// 216. 组合总和 III
// 找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。

// 说明：
// 所有数字都是正整数。
// 解集不能包含重复的组合。

var res [][]int

func combinationSum3(k int, n int) [][]int {
	res = make([][]int, 0)
	path := make([]int, 0)
	solve(k, n, 1, path)
	return res
}

func solve(k, n, start int, path []int) {
	if n < 0 {
		return
	}
	if n == 0 && len(path) == k {
		tmp := make([]int, k)
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := start; i <= 9; i++ {
		path = append(path, i)
		solve(k, n-i, i+1, path)
		path = path[:len(path)-1]
	}
}
