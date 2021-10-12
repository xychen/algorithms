package problem0090

// 90. 子集 II
// 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := make([][]int, 0)
	path := make([]int, 0)
	visted := make([]bool, n)

	var solve func(start int, path []int)
	solve = func(start int, path []int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		for i := start; i < n; i++ {
			if visted[i] || (i > 0 && nums[i] == nums[i-1] && !visted[i-1]) {
				continue
			}
			visted[i] = true
			path = append(path, nums[i])
			solve(i+1, path)

			path = path[:len(path)-1]
			visted[i] = false
		}
	}
	solve(0, path)
	return res
}
