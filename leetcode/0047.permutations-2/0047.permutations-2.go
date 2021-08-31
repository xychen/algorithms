package problem0047

var res [][]int
var vis []bool

func permuteUnique(nums []int) [][]int {
	res = make([][]int, 0)
	trace := make([]int, 0)
	vis = make([]bool, len(nums))
	backtrack(nums, trace)
	return res
}

func backtrack(nums []int, trace []int) {
	if len(nums) == len(trace) {
		tmp := make([]int, len(trace))
		copy(tmp, trace)
		res = append(res, tmp)
		return
	}

	for i, num := range nums {
		//（已经走过了 或者 （每次填入[当前位置]的数一定是这个数所在重复数集合中「从左往右第一个未被填过的数字」）
		// 1.重复组合 （num == nums[i-1]）  2.从左往右第一个未被填过的数字（！vis[i-1], i-1未填过，i肯定不是第一个）
		if vis[i] || (i > 0 && num == nums[i-1] && !vis[i-1]) {
			continue
		}
		//做选择
		trace = append(trace, num)
		vis[i] = true
		//进入下一层决策树
		backtrack(nums, trace)
		//取消选择
		trace = trace[:len(trace)-1]
		vis[i] = false
	}
}
