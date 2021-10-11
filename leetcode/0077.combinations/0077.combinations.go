package problem0077

// 77. 组合
// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
// 你可以按 任何顺序 返回答案。

var res [][]int

func combine(n int, k int) [][]int {
	res = make([][]int, 0)
	path := make([]int, 0)
	solve(n, k, 1, path)
	return res
}

func solve(n, k, start int, path []int) {
	if k == len(path) {
		tmp := make([]int, k)
		copy(tmp, path)
		res = append(res, tmp)
	}
	if start > n {
		return
	}
	for i := start; i <= n; i++ {
		path = append(path, i)
		solve(n, k, i+1, path)
		path = path[:len(path)-1]
	}
}
