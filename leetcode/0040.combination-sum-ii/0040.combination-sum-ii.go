package problem0040

// 40. 组合总和 II
// 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的每个数字在每个组合中只能使用一次。
// 注意：解集不能包含重复的组合。

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) <= 0 {
		return [][]int{}
	}
	res := make([][]int, 0)
	sort.Ints(candidates)
	// 下标 => 是否被访问
	visited := make([]bool, len(candidates))
	var backtrack func(target, startIndex int, path []int)
	backtrack = func(target, startIndex int, path []int) {
		if target == 0 {
			tmpPath := make([]int, len(path))
			copy(tmpPath, path)
			res = append(res, tmpPath)
			return
		}

		for i := startIndex; i < len(candidates); i++ {
			// 剪枝
			if target-candidates[i] < 0 {
				break
			}
			// 同一层中，相同的数，保证左边的优先访问到
			if i > 0 && candidates[i] == candidates[i-1] && !visited[i-1] {
				continue
			}
			visited[i] = true
			path = append(path, candidates[i])
			backtrack(target-candidates[i], i+1, path)
			path = path[:len(path)-1]
			visited[i] = false
		}
	}
	path := make([]int, 0)
	backtrack(target, 0, path)
	return res
}
