package problem0078

// 78. 子集
// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

func subsets(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0)
	path := make([]int, 0)

	var solve func(start int, path []int)
	solve = func(start int, path []int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		if start >= n {
			return
		}

		for i := start; i < n; i++ {
			path = append(path, nums[i])
			solve(i+1, path)
			path = path[:len(path)-1]
		}
	}
	solve(0, path)
	return res
}
