package problem62

// 剑指 Offer 62. 圆圈中最后剩下的数字
// 约瑟夫环问题

func lastRemaining(n int, m int) int {
	if n < 1 || m < 1 {
		return -1
	}

	last := 0
	for i := 2; i <= n; i++ {
		last = (last + m) % i
	}
	return last
}
