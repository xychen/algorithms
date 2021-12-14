package problem12

// 79. 单词搜索
// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	var backtrace func(board [][]byte, word string, cur, i, j int) bool
	backtrace = func(board [][]byte, word string, cur, i, j int) bool {
		if cur >= len(word) {
			return true
		}
		// 越界
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
			return false
		}
		if board[i][j] != byte(word[cur]) {
			return false
		}
		if visited[i][j] {
			return false
		}
		// 做选择
		visited[i][j] = true

		if backtrace(board, word, cur+1, i-1, j) || backtrace(board, word, cur+1, i+1, j) || backtrace(board, word, cur+1, i, j-1) || backtrace(board, word, cur+1, i, j+1) {
			return true
		}
		// 回退
		visited[i][j] = false
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if backtrace(board, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}
