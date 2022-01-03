package problem0051

//n皇后问题
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)

	var backtrack func(row int, selected [][]byte)
	backtrack = func(row int, selected [][]byte) {
		if row >= n {
			tmplist := make([]string, n)
			for i := 0; i < n; i++ {
				tmplist[i] = string(selected[i])
			}
			res = append(res, tmplist)
			return
		}

		for col := 0; col < n; col++ {
			if !isValid(row, col, selected) {
				continue
			}
			selected[row][col] = 'Q'
			backtrack(row+1, selected)
			selected[row][col] = '.'
		}
	}

	selected := make([][]byte, n)
	for i := 0; i < n; i++ {
		selected[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			selected[i][j] = '.'
		}
	}
	backtrack(0, selected)
	return res
}

func isValid(row, col int, matrix [][]byte) bool {
	n := len(matrix)
	// 竖方向
	for i := 0; i < row; i++ {
		if matrix[i][col] == 'Q' {
			return false
		}
	}

	// 左斜上方
	i, j := row-1, col-1
	for i >= 0 && j >= 0 {
		if matrix[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}

	// 右斜上方
	i, j = row-1, col+1
	for i >= 0 && j < n {
		if matrix[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}
	return true
}
