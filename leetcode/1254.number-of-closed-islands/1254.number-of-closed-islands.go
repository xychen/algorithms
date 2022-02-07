package problem1254

// 1254. 统计封闭岛屿的数目
// 二维矩阵 grid 由 0 （土地）和 1 （水）组成。岛是由最大的4个方向连通的 0 组成的群，封闭岛是一个 完全 由1包围（左、上、右、下）的岛。

// 请返回 封闭岛屿 的数目。

func closedIsland(grid [][]int) int {
	var dfs func(grid [][]int, i, j int)
	dfs = func(grid [][]int, i, j int) {
		n, m := len(grid), len(grid[0])
		if i < 0 || i >= n || j < 0 || j >= m {
			return
		}
		// 非岛屿
		if grid[i][j] == 1 {
			return
		}
		// 水淹
		grid[i][j] = 1
		dfs(grid, i-1, j)
		dfs(grid, i+1, j)
		dfs(grid, i, j-1)
		dfs(grid, i, j+1)
	}
	n, m := len(grid), len(grid[0])
	// 先把边界淹掉
	for i := 0; i < n; i++ {
		dfs(grid, i, 0)
		dfs(grid, i, m-1)
	}
	for j := 0; j < m; j++ {
		dfs(grid, 0, j)
		dfs(grid, n-1, j)
	}

	res := 0
	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			if grid[i][j] == 0 {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}
