package problem0046

//全排列，优化是否遍历过
var res2 [][]int
var vis []bool

func permute2(nums []int) [][]int {
	res = make([][]int, 0)
	trace := make([]int, 0)
	vis = make([]bool, len(nums))
	backtrack2(nums, trace)
	return res
}

func backtrack2(nums []int, trace []int) {
	if len(nums) == len(trace) {
		tmp := make([]int, len(trace))
		copy(tmp, trace)
		res = append(res, tmp)
		return
	}

	for i, num := range nums {
		//已经走过了
		if vis[i] {
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
