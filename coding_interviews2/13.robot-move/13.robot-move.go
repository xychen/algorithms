package problem13

// 剑指 Offer 13. 机器人的运动范围
// 地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

func movingCount(m int, n int, k int) int {
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var backtrace func(i, j int) int
	backtrace = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n {
			return 0
		}
		if bitSum(i)+bitSum(j) > k || visited[i][j] {
			return 0
		}
		// 当前节点可以访问
		sum := 1
		visited[i][j] = true
		// 加上他邻居可以访问的节点
		sel := [][]int{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
		for _, item := range sel {
			sum += backtrace(item[0], item[1])
		}
		return sum
	}
	return backtrace(0, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func bitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n = n / 10
	}
	return sum
}
