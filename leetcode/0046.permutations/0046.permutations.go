package problem0046

//全排列

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	visited := make([]bool, len(nums))
	var backtrack func(nums, path []int)
	backtrack = func(nums, path []int) {
		if len(path) == len(nums) {
			tmpPath := make([]int, len(path))
			copy(tmpPath, path)
			res = append(res, tmpPath)
			return
		}

		for i, num := range nums {
			if visited[i] {
				continue
			}
			path = append(path, num)
			visited[i] = true
			backtrack(nums, path)
			path = path[:len(path)-1]
			visited[i] = false
		}
	}
	path := make([]int, 0, len(nums))
	backtrack(nums, path)
	return res
}
