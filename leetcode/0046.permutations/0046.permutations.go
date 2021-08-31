package problem0046

//全排列

var res [][]int

func permute(nums []int) [][]int {
	res = make([][]int, 0)
	trace := make([]int, 0)
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

	for _, num := range nums {
		//已经走过了
		if contain(trace, num) {
			continue
		}
		//做选择
		trace = append(trace, num)
		//进入下一层决策树
		backtrack(nums, trace)
		//取消选择
		trace = trace[:len(trace)-1]
	}
}

func contain(arr []int, target int) bool {
	for _, a := range arr {
		if a == target {
			return true
		}
	}
	return false
}
