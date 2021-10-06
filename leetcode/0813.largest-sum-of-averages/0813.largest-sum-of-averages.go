package problem0813

// 813. 最大平均值和的分组
// 我们将给定的数组 A 分成 K 个相邻的非空子数组 ，我们的分数由每个子数组内的平均值的总和构成。计算我们所能得到的最大分数是多少。

// 注意我们必须使用 A 数组中的每一个数进行分组，并且分数不一定需要是整数。

// 解法一： 动归+压缩数组
func largestSumOfAverages(nums []int, k int) float64 {
	n := len(nums)
	sums := make([]int, n+1)
	for i := 0; i < n; i++ {
		sums[i+1] = sums[i] + nums[i]
	}
	// dp[i] 表示nums[i:]元素分成k组的最大分数
	// dp(i, k) = max(dp(j, k - 1) + average(i, j - 1))， j > i
	dp := make([]float64, n)
	// k = 1时，dp[i]为平均值
	for i := 0; i < n; i++ {
		dp[i] = float64(sums[n]-sums[i]) / float64(n-i)
	}
	K := k
	for k = 2; k <= K; k++ {
		for i := 0; i < n; i++ {
			// maxSingleGroupCount := (n - i) - k + 1
			// i + maxSingleGroupCount => n - k + 1
			for j := i + 1; j <= n-k+1; j++ {
				dp[i] = max(dp[i], dp[j]+float64(sums[j]-sums[i])/float64(j-i))
			}
		}
	}
	return dp[0]
}

// 解法二： 动归+dp全数组
func largestSumOfAverages2(nums []int, k int) float64 {
	n := len(nums)
	sums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	if k == 1 {
		return float64(sums[n]) / float64(n)
	}

	// dp[i][k] 表示前i个元素，分成k个相邻的非空子数组（最终答案为dp[n][k]）
	dp := make([][]float64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]float64, k+1)
		dp[i][0] = float64(sums[i])
		dp[i][1] = average(sums, 1, i)
	}

	// dp(i, k) = max(dp(j, k - 1) + average(j + 1, i))
	// dp(i, 0) = average(0, i)
	for i := 1; i <= n; i++ {
		for k1 := 2; k1 <= k; k1++ {
			maxSingleGroupCount := i - k1
			for j := i - maxSingleGroupCount; j <= i; j++ {
				tmp := dp[j-1][k1-1] + average(sums, j, i)
				dp[i][k1] = max(dp[i][k1], tmp)
			}
		}
	}
	return dp[n][k]
}

// [left, right]的平均值, left从1开始
func average(sums []int, left, right int) float64 {
	return float64(sums[right]-sums[left-1]) / float64(right-(left-1))
}

// 解法三： 回溯
func solve(nums []int, k int) float64 {
	// fmt.Println(nums, k)
	n := len(nums)
	if n < k {
		return float64(0)
	}
	if n == k {
		return sum(nums)
	}
	if k == 1 {
		return avg(nums)
	}
	res := float64(0)
	// 分组内最大的元素数量为： len(nums) - (k-1)
	singleMaxCount := n - (k - 1)
	for i := 1; i <= singleMaxCount; i++ {
		if n-i <= 0 {
			break
		}
		if n-i < k-1 {
			break
		}
		//做选择
		tmp := solve(nums[:n-i], k-1) + avg(nums[n-i:])
		// fmt.Println(res, tmp)
		res = max(res, tmp)
		//回退
	}
	return res
}

func sum(nums []int) float64 {
	s := 0
	for _, num := range nums {
		s += num
	}
	return float64(s)
}

func avg(nums []int) float64 {
	return sum(nums) / float64(len(nums))
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
