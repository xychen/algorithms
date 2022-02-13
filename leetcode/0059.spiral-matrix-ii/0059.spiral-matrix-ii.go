package problem0059

// 59. 螺旋矩阵 II
// 给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。

// 输入：n = 3
// 输出：[[1,2,3],[8,9,4],[7,6,5]]

// 左闭右开，所有循环都按照同一个约定
func generateMatrix(n int) [][]int {
	loop := n / 2
	startx, starty := 0, 0
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	k := 1
	for loop > 0 {
		i, j := startx, starty
		for ; j < n-starty-1; j++ {
			res[i][j] = k
			k += 1
		}
		for ; i < n-startx-1; i++ {
			res[i][j] = k
			k += 1
		}
		for ; j > starty; j-- {
			res[i][j] = k
			k += 1
		}
		for ; i > startx; i-- {
			res[i][j] = k
			k += 1
		}

		startx += 1
		starty += 1
		loop -= 1
	}
	if n%2 == 1 {
		res[n/2][n/2] = k
	}
	return res
}
