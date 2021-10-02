package problem1043

// 分隔数组以得到最大和

// 给你一个整数数组 arr，请你将该数组分隔为长度最多为 k 的一些（连续）子数组。分隔完成后，每个子数组的中的所有值都会变为该子数组中的最大值。
// 返回将数组分隔变换后能够得到的元素最大和。
// 注意，原数组和分隔后的数组对应顺序应当一致，也就是说，你只能选择分隔数组的位置而不能调整数组中的顺序。

// 方法一：动态规划
func maxSumAfterPartitioning(arr []int, k int) int {

	// dp[i] 表示arr[i]结束的 分隔数组最大和
	n := len(arr)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		curMax := -1
		for j := 1; j <= k; j++ {
			index := i - j
			if index < 0 {
				break
			}
			curMax = max(curMax, arr[index])
			dp[i] = max(dp[i], dp[index]+curMax*j)
		}
	}
	return dp[n]
}

// 方法二： 回溯，会超时
func solve(arr []int, k int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	maxRes := 0
	for i := 1; i <= k; i++ {
		//做选择
		index := len(arr) - i
		if index < 0 {
			break
		}
		// 数组左闭右开
		tmp := solve(arr[:index], k) + maxValue(arr[index:])*i
		maxRes = max(maxRes, tmp)
		//回退
	}
	return maxRes
}

func maxValue(nums []int) int {
	maxVal := nums[0]
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}
	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
