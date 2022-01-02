package problem0047

import "sort"

// 47. 全排列 II
// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	visited := make([]bool, len(nums))
	var backtrack func(nums, path []int)
	backtrack = func(nums, path []int) {
		if len(nums) == len(path) {
			tmpPath := make([]int, len(path))
			copy(tmpPath, path)
			res = append(res, tmpPath)
			return
		}

		for i, num := range nums {
			if visited[i] || (i > 0 && num == nums[i-1] && !visited[i-1]) {
				continue
			}
			visited[i] = true
			path = append(path, num)
			backtrack(nums, path)
			path = path[:len(path)-1]
			visited[i] = false
		}
	}
	path := make([]int, 0, len(nums))
	backtrack(nums, path)
	return res
}
