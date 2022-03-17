package problem66

// 剑指 Offer 66. 构建乘积数组

func constructArr(a []int) []int {
	if len(a) < 1 {
		return []int{}
	}
	n := len(a)
	b := make([]int, n)
	c := make([]int, n)
	d := make([]int, n)
	c[0] = 1
	for i := 1; i < n; i++ {
		c[i] = c[i-1] * a[i-1]
	}

	d[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		d[i] = d[i+1] * a[i+1]
	}

	b[0] = d[0]
	b[n-1] = c[n-1]

	for i := 1; i < n-1; i++ {
		b[i] = c[i] * d[i]
	}
	return b
}
