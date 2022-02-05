package problem0053

// 贪心解法
func maxSubArrayGreedy(nums []int) int {
	res, sum := -10000000, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > res {
			res = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return res
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
