package problem0494

// 给你一个整数数组 nums 和一个整数 target 。
// 向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
// 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
// 返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目

// 方法一：转成动态规划的0-1背包问题
func findTargetSumWays(nums []int, target int) int {
	//转成背包问题
	n, sum := len(nums), 0
	for _, num := range nums {
		sum += num
	}
	// (sum - neg) - neg = target =>  neg = (sum - target) / 2
	// 挑选出一些元素，和等于neg
	if sum < target {
		return 0
	}
	if (sum-target)%2 == 1 {
		return 0
	}
	m := (sum - target) / 2
	// dp[i][j] => 前i个元素，和等于j的数目
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	// ***** 注意base case 和 416题的区别，太烧脑了  ******
	dp[0][0] = 1

	for i := 1; i <= n; i++ {
		v := nums[i-1]
		// 因为neg（m的值）可以是0，所以从0开始遍历
		for j := 0; j <= m; j++ {
			if j >= v {
				//选v
				dp[i][j] += dp[i-1][j-v]
				//不选v
				dp[i][j] += dp[i-1][j]
			} else {
				dp[i][j] += dp[i-1][j]
			}
		}
	}
	// fmt.Println(m, dp)
	return dp[n][m]
}

// 方法二： 回溯

func findTargetSumWays2(nums []int, target int) int {
	return solve(nums, target)
}

func solve(nums []int, target int) int {
	n := len(nums)
	if target == 0 && n == 0 {
		return 1
	}
	if n == 0 {
		return 0
	}
	res := 0
	//做选择
	res += solve(nums[:n-1], target-nums[n-1])
	//回退
	//做选择
	res += solve(nums[:n-1], target+nums[n-1])
	//回退
	return res
}
