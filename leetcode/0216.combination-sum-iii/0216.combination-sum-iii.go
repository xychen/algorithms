package problem0216

// 216. 组合总和 III
// 找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。

// 说明：
// 所有数字都是正整数。
// 解集不能包含重复的组合。

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	var backtrack func(k, sum, statIndex int, path []int)
	backtrack = func(k, sum, statIndex int, path []int) {
		if k == len(path) {
			if sum == 0 {
				tmpPath := make([]int, k)
				copy(tmpPath, path)
				res = append(res, tmpPath)
			}
			return
		}

		// 9-(k-len(path))+1为剪枝操作
		for i := statIndex; i <= 9-(k-len(path))+1; i++ {
			// 做选择
			path = append(path, i)
			backtrack(k, sum-i, i+1, path)
			// 回退
			path = path[:len(path)-1]
		}
	}
	path := make([]int, 0)
	backtrack(k, n, 1, path)
	return res
}
