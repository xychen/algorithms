package problem0059

// 59. 螺旋矩阵 II
// 给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。

// 输入：n = 3
// 输出：[[1,2,3],[8,9,4],[7,6,5]]

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	curNum := 1
	for i := n; i > 0; i -= 2 {
		curNum = handle(n, i, curNum, res)
	}
	return res
}

func handle(N, n, start int, nums [][]int) int {
	curNum := start
	span := (N - n) / 2
	// 上横
	for j := 0; j < n; j++ {
		nums[span][j+span] = curNum
		curNum++
	}
	if n == 1 {
		return curNum
	}
	// 右竖(短)
	for i := 1; i < n-1; i++ {
		nums[span+i][span+n-1] = curNum
		curNum++
	}

	// 下横
	for j := n - 1; j >= 0; j-- {
		nums[span+n-1][j+span] = curNum
		curNum++
	}

	// 左竖（短）
	for i := n - 2; i > 0; i-- {
		nums[span+i][span] = curNum
		curNum++
	}

	return curNum
}
