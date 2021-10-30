package problem0064

// 64. 最小路径和
// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

// 说明：每次只能向下或者向右移动一步。

func minPathSum(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
		if i == 0 {
			res[0][0] = grid[0][0]
		} else {
			res[i][0] = res[i-1][0] + grid[i][0]
		}
	}
	for j := 1; j < m; j++ {
		res[0][j] = res[0][j-1] + grid[0][j]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			res[i][j] = min(res[i-1][j], res[i][j-1]) + grid[i][j]
		}
	}
	return res[n-1][m-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
