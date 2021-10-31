package problem04

// 剑指 Offer 04. 二维数组中的查找
// 在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	col := len(matrix[0])
	row := len(matrix)
	r, c := 0, col-1
	for r < row && c >= 0 {
		if matrix[r][c] == target {
			return true
		} else if matrix[r][c] < target {
			r++
		} else if matrix[r][c] > target {
			c--
		}
	}
	return false
}
