package problem0078

// 78. 子集
// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	var backtrack func(nums []int, startIndex int, path []int)
	backtrack = func(nums []int, startIndex int, path []int) {
		tmpPath := make([]int, len(path))
		copy(tmpPath, path)
		res = append(res, tmpPath)

		for i := startIndex; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(nums, i+1, path)
			path = path[:len(path)-1]
		}
	}
	path := make([]int, 0)
	backtrack(nums, 0, path)
	return res
}
