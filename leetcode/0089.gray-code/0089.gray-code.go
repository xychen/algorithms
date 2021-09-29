package problem0089

// 格雷编码

func grayCode(n int) []int {
	res := []int{0}
	head := 1
	for i := 1; i <= n; i++ {
		for j := len(res) - 1; j >= 0; j-- {
			res = append(res, head+res[j])
		}
		head = head << 1
	}
	return res
}
