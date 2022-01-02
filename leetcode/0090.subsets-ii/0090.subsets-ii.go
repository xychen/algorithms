package problem0090

// 90. 子集 II
// 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	visited := make([]bool, len(nums))
	var backtrack func(nums []int, startIndex int, path []int)
	backtrack = func(nums []int, startIndex int, path []int) {
		tmpPath := make([]int, len(path))
		copy(tmpPath, path)
		res = append(res, tmpPath)

		for i := startIndex; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
				continue
			}
			visited[i] = true
			path = append(path, nums[i])
			backtrack(nums, i+1, path)
			visited[i] = false
			path = path[:len(path)-1]
		}
	}
	path := make([]int, 0)
	backtrack(nums, 0, path)
	return res
}
