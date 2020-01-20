package problem003

//二位数组查找元素

func Find(matrix []int, rows, columns, target int) bool {
	if matrix == nil || rows <= 0 || columns <= 0 {
		return false
	}
	row := 0
	column := columns - 1
	for row < rows && column >= 0 {
		index := row*columns + column
		if matrix[index] == target {
			return true
		} else if matrix[index] > target {
			column -= 1
		} else {
			row += 1
		}
	}

	return false
}
