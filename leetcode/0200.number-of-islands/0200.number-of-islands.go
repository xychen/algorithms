package problem0200

// 200. 岛屿数量
// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
// 此外，你可以假设该网格的四条边均被水包围。

func numIslands(grid [][]byte) int {
	var dfs func(grid [][]byte, i, j int)
	dfs = func(grid [][]byte, i, j int) {
		n, m := len(grid), len(grid[0])
		if i < 0 || i >= n || j < 0 || j >= m {
			return
		}
		if grid[i][j] == '0' {
			return
		}
		// 水淹
		grid[i][j] = '0'
		dfs(grid, i-1, j)
		dfs(grid, i+1, j)
		dfs(grid, i, j+1)
		dfs(grid, i, j-1)
	}
	n, m := len(grid), len(grid[0])
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}
