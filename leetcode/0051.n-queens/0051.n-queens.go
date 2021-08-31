package problem0051

import "strings"

//n皇后问题
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]

var res [][]string

func solveNQueens(n int) [][]string {
	res = make([][]string, 0)
	//trace初始化
	trace := make([][]string, n)
	for i := 0; i < n; i++ {
		trace[i] = make([]string, n)
		for j := 0; j < n; j++ {
			trace[i][j] = "."
		}
	}
	backtrack(0, n, trace)
	return res
}

func backtrack(row, n int, trace [][]string) {
	if n == row {
		tmp := make([]string, n)
		for i, items := range trace {
			tmp[i] = strings.Join(items, "")
		}
		res = append(res, tmp)
		return
	}

	for column := 0; column < n; column++ {
		if !isValid(row, column, trace) {
			continue
		}
		//选择
		trace[row][column] = "Q"
		backtrack(row+1, n, trace)
		//回退
		trace[row][column] = "."
	}
}

func isValid(row, col int, trace [][]string) bool {
	//左上是否有
	for i := 1; (row-i >= 0) && (col-i >= 0); i++ {
		if trace[row-i][col-i] == "Q" {
			return false
		}
	}
	//同一列中是否有
	for i := 1; row-i >= 0; i++ {
		if trace[row-i][col] == "Q" {
			return false
		}
	}
	//右上是否有
	for i := 1; (row-i >= 0) && (col+i) < len(trace); i++ {
		if trace[row-i][col+i] == "Q" {
			return false
		}
	}
	return true
}
