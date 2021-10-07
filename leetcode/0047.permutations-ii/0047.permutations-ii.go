package problem0047

// 47. 全排列 II
// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

import "sort"

var res [][]int
var visted map[int]bool

func permuteUnique(nums []int) [][]int {
	res = make([][]int, 0)
	visted = make(map[int]bool)
	path := make([]int, 0)
	sort.Ints(nums)
	solve(nums, path)
	return res
}
func solve(nums []int, path []int) {
	if len(nums) == len(path) {
		tmp := make([]int, len(nums))
		copy(tmp, path)
		res = append(res, tmp)
	}
	for i := 0; i < len(nums); i++ {
		if v, ok := visted[i]; ok && v {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && visted[i-1] == false {
			continue
		}
		path = append(path, nums[i])
		visted[i] = true
		solve(nums, path)

		visted[i] = false
		path = path[:len(path)-1]
	}
}
