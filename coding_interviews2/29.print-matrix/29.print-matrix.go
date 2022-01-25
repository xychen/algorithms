package problem29

// 剑指 Offer 29. 顺时针打印矩阵

var res []int
var c int

func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	if n == 0 {
		return []int{}
	}
	m := len(matrix[0])
	res = make([]int, n*m)
	c = 0
	start := 0
	for 2*start < n && 2*start < m {
		print(matrix, start)
		start++
	}
	return res
}

func print(matrix [][]int, start int) {
	n := len(matrix)
	m := len(matrix[0])
	// fmt.Println(start)
	// 长上行
	for j := start; j < m-start; j++ {
		res[c] = matrix[start][j]
		c++
	}
	// 短右列
	for i := start + 1; i < n-start-1; i++ {
		res[c] = matrix[i][m-start-1]
		c++
	}
	if start != n-start-1 {
		// 长下行
		for j := m - start - 1; j >= start; j-- {
			res[c] = matrix[n-start-1][j]
			c++
		}
	}

	if start != m-start-1 {
		// 短左列
		for i := n - start - 2; i > start; i-- {
			res[c] = matrix[i][start]
			c++
		}
	}
}
