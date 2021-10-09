package problem0040

// 40. 组合总和 II
// 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的每个数字在每个组合中只能使用一次。
// 注意：解集不能包含重复的组合。

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	n := len(candidates)
	res := make([][]int, 0)
	path := make([]int, 0)
	visited := make([]bool, n)
	sort.Ints(candidates)
	var solve func(target, start int, path []int)
	solve = func(target, start int, path []int) {
		if target == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		if target < 0 {
			return
		}
		for i := start; i < n; i++ {
			if visited[i] || (i > 0 && candidates[i] == candidates[i-1] && !visited[i-1]) {
				continue
			}
			path = append(path, candidates[i])
			visited[i] = true
			solve(target-candidates[i], i+1, path)

			visited[i] = false
			path = path[:len(path)-1]
		}
	}

	solve(target, 0, path)
	return res
}
