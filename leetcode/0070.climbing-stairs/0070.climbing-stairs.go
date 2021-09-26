package problem0070

// 爬楼梯

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	pre, cur := 1, 2
	var res int
	for i := 3; i <= n; i++ {
		res = pre + cur
		pre, cur = cur, res
	}
	return res
}
