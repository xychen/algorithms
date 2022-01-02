package problem0491

func findSubsequences(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{}
	}
	res := make([][]int, 0)
	var backtrack func(nums []int, startIndex int, path []int)
	backtrack = func(nums []int, startIndex int, path []int) {
		if len(path) >= 2 {
			tmpPath := make([]int, len(path))
			copy(tmpPath, path)
			res = append(res, tmpPath)
		}
		// 本层去重
		visited := make(map[int]bool)
		for i := startIndex; i < len(nums); i++ {
			// 非递增数字
			if len(path) > 0 && nums[i] < path[len(path)-1] {
				continue
			}
			// 处理本层相同元素的问题
			if _, ok := visited[nums[i]]; ok {
				continue
			}

			visited[nums[i]] = true
			path = append(path, nums[i])
			backtrack(nums, i+1, path)
			path = path[:len(path)-1]
		}
	}
	path := make([]int, 0)
	backtrack(nums, 0, path)
	return res
}
