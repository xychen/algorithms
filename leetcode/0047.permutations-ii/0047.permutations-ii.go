package problem0047

// 47. 全排列 II
// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

import "sort"

func permuteUnique(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	res := make([][]int, 0)
	visted := make([]bool, n)
	path := make([]int, 0)
	var solve func(path []int)
	solve = func(path []int) {
		if len(nums) == len(path) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			// 要解决重复问题，我们只要设定一个规则，保证在填第idx 个数的时候重复数字只会被填入一次即可。
			// 而在本题解中，我们选择对原数组排序，保证相同的数字都相邻，然后每次填入的数一定是这个数所在重复数集合中「从左往右第一个未被填过的数字」，
			// 即如下的判断条件：i > 0 && nums[i] == nums[i-1] && !visted[i-1]  , !visted[i-1] 表示i-1没使用，那i肯定不是第一个
			if visted[i] || (i > 0 && nums[i] == nums[i-1] && !visted[i-1]) {
				continue
			}
			visted[i] = true
			path = append(path, nums[i])
			solve(path)
			path = path[:len(path)-1]
			visted[i] = false
		}
	}
	solve(path)
	return res
}
