package problem0077

// 77. 组合
// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
// 你可以按 任何顺序 返回答案。

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	var backtrack func(n, k, startIndex int, path []int)
	backtrack = func(n, k, startIndex int, path []int) {
		if len(path) == k {
			tmpPath := make([]int, k)
			copy(tmpPath, path)
			res = append(res, tmpPath)
			return
		}
		// n-(k-len(path))+1 为剪枝操作
		for i := startIndex; i <= n-(k-len(path))+1; i++ {
			// 做选择
			path = append(path, i)
			backtrack(n, k, i+1, path)
			// 回退
			path = path[:len(path)-1]
		}
	}
	path := make([]int, 0)
	backtrack(n, k, 1, path)
	return res
}
