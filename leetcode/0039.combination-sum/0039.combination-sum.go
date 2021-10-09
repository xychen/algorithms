package problem0039

// 39. 组合总和
// 给定一个无重复元素的正整数数组 candidates 和一个正整数 target ，找出 candidates 中所有可以使数字和为目标数 target 的唯一组合。
// candidates 中的数字可以无限制重复被选取。如果至少一个所选数字数量不同，则两种组合是唯一的。
// 对于给定的输入，保证和为 target 的唯一组合数少于 150 个。

var res [][]int

func combinationSum(candidates []int, target int) [][]int {
	res = make([][]int, 0)
	path := make([]int, 0)
	solve(candidates, target, 0, path)
	return res
}

func solve(candidates []int, target int, start int, path []int) {
	if target == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	if target < 0 {
		return
	}
	for i := start; i < len(candidates); i++ {
		path = append(path, candidates[i])
		solve(candidates, target-candidates[i], i, path)
		path = path[:len(path)-1]
	}
}
