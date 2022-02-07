package problem1905

// 1905. 统计子岛屿
// 给你两个 m x n 的二进制矩阵 grid1 和 grid2 ，它们只包含 0 （表示水域）和 1 （表示陆地）。一个 岛屿 是由 四个方向 （水平或者竖直）上相邻的 1 组成的区域。任何矩阵以外的区域都视为水域。
// 如果 grid2 的一个岛屿，被 grid1 的一个岛屿 完全 包含，也就是说 grid2 中该岛屿的每一个格子都被 grid1 中同一个岛屿完全包含，那么我们称 grid2 中的这个岛屿为 子岛屿 。
// 请你返回 grid2 中 子岛屿 的 数目 。

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	var dfs func(grid [][]int, i, j int)
	dfs = func(grid [][]int, i, j int) {
		n, m := len(grid), len(grid[0])
		if i < 0 || i >= n || j < 0 || j >= m {
			return
		}
		if grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		dfs(grid, i-1, j)
		dfs(grid, i+1, j)
		dfs(grid, i, j-1)
		dfs(grid, i, j+1)
	}
	n, m := len(grid1), len(grid1[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// 肯定不是子岛屿，淹掉
			if grid2[i][j] == 1 && grid1[i][j] == 0 {
				dfs(grid2, i, j)
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid2[i][j] == 1 {
				res++
				dfs(grid2, i, j)
			}
		}
	}
	return res
}
