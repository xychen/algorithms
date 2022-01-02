package problem0039

import "sort"

// 39. 组合总和
// 给定一个无重复元素的正整数数组 candidates 和一个正整数 target ，找出 candidates 中所有可以使数字和为目标数 target 的唯一组合。
// candidates 中的数字可以无限制重复被选取。如果至少一个所选数字数量不同，则两种组合是唯一的。
// 对于给定的输入，保证和为 target 的唯一组合数少于 150 个。

// 剪枝解法
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) <= 0 {
		return [][]int{}
	}
	res := make([][]int, 0)
	// 剪枝操作需要进行排序
	sort.Ints(candidates)
	var backtrack func(candidates []int, target, startIndex int, path []int)
	backtrack = func(candidates []int, target, startIndex int, path []int) {
		if target == 0 {
			tmpPath := make([]int, len(path))
			copy(tmpPath, path)
			res = append(res, tmpPath)
			return
		}

		for i := startIndex; i < len(candidates); i++ {
			// 后边的值肯定都会出现小于0，不必再进入下一层（剪枝操作）
			if target-candidates[i] < 0 {
				break
			}
			path = append(path, candidates[i])
			backtrack(candidates, target-candidates[i], i, path)
			path = path[:len(path)-1]
		}
	}
	path := make([]int, 0)
	backtrack(candidates, target, 0, path)
	return res
}

// 普通解法
func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) <= 0 {
		return [][]int{}
	}
	res := make([][]int, 0)
	var backtrack func(candidates []int, target, startIndex int, path []int)
	backtrack = func(candidates []int, target, startIndex int, path []int) {
		if target <= 0 {
			if target == 0 {
				tmpPath := make([]int, len(path))
				copy(tmpPath, path)
				res = append(res, tmpPath)
			}
			return
		}

		for i := startIndex; i < len(candidates); i++ {
			path = append(path, candidates[i])
			backtrack(candidates, target-candidates[i], i, path)
			path = path[:len(path)-1]
		}
	}
	path := make([]int, 0)
	backtrack(candidates, target, 0, path)
	return res
}
