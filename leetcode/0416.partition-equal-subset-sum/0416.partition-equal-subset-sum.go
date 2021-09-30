package problem0416

// 分割等和子集
// 给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

// 方法一：dp+压缩空间
func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	sum, maxVal := 0, 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}
	if sum%2 == 1 {
		return false
	}
	sum = sum / 2
	if maxVal > sum {
		return false
	}
	// dp[i][j] 的意思是nums前i个元素，是否存在相加等于j的情况
	// i等于0的时候，表示前0个元素，肯定是false
	dp := make([]bool, sum+1)
	dp[nums[0]] = true
	for i := 1; i < n; i++ {
		v := nums[i]
		for j := sum; j >= 1; j-- {
			// if v <= j { //不选 或 选 有一个ok就为true
			// 	dp[j] = dp[j] || dp[j-v]
			// }
			if j == v {
				dp[j] = true
			} else if j > v {
				dp[j] = dp[j] || dp[j-v]
			}
		}
	}
	return dp[sum]
}

// 方法二：动态规划 （未压缩空间）

func canPartition2(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	sum, maxVal := 0, 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}
	if sum%2 == 1 {
		return false
	}
	sum = sum / 2
	if maxVal > sum {
		return false
	}
	// dp[i][j] 的意思是nums前i个元素，是否存在相加等于j的情况
	dp := make([][]bool, n)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, sum+1)
	}

	/**
	* 边界情况
	* 因为 【注意1】【注意2】，为了不特殊处理，可以通过增加一行，同时第0列设置为true即可
	**/

	//注意1：优先处理好第0行，防止出现 i-1 等于负数的情况
	dp[0][nums[0]] = true
	for i := 1; i <= n; i++ {
		v := nums[i]
		for j := 1; j <= sum; j++ {
			// 注意2：如果想把j==v 和 j > v 的情况合并，就需要特殊梳理一下 j=0的值（提前初始化成true）
			if j == v {
				//只选当前这个肯定是ok的
				dp[i][j] = true
			} else if j > v {
				//不选 或 选 有一个ok就为true
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else { //不能选了
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][sum]
}

// 方法三： 回溯 + 记忆化搜索
var dp map[int]bool
var sel []int

func canPartition3(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 == 1 {
		return false
	}
	sum = sum / 2
	dp = make(map[int]bool)
	sel = make([]int, len(nums))
	return solve(nums, sum)
}

func solve(nums []int, target int) bool {
	if target == 0 {
		return true
	}
	if target < 0 {
		return false
	}

	if v, ok := dp[target]; ok {
		return v
	}
	var res bool
	for i, n := range nums {
		if 1 == sel[i] {
			continue
		}
		//做选择
		sel[i] = 1
		res = solve(nums, target-n)
		if res == true {
			break
		}
		//回退
		sel[i] = 0
	}
	dp[target] = res
	return res
}
