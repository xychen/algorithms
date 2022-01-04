package problem0509

// 509. 斐波那契数

func fib(n int) int {
	if n < 2 {
		return n
	}
	pre, cur := 0, 1
	for i := 2; i <= n; i++ {
		tmp := pre + cur
		pre, cur = cur, tmp
	}
	return cur
}
