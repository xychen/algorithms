package problem0198

var table []int

//递归 + 记忆化搜索
func rob(nums []int) int {
	table = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		table[i] = -1
	}
	return slove(nums, len(nums)-1)
}

func slove(nums []int, i int) int {
	if i < 0 {
		return 0
	}
	if table[i] >= 0 {
		return table[i]
	}
	table[i] = max(nums[i]+slove(nums, i-2), slove(nums, i-1))
	return table[i]
}

//自底向上
func rob2(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	table = make([]int, len(nums))
	table[0] = nums[0]
	table[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		table[i] = max(nums[i]+table[i-2], table[i-1])
	}
	return table[len(nums)-1]
}

//自底向上: 去冗余空间（滚动数组）
func rob3(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	f1 := nums[0]
	f2 := max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		tmp := max(nums[i]+f1, f2)
		f1 = f2
		f2 = tmp
	}
	return f2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
