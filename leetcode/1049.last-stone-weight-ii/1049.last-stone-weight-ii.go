package problem1049

// 1049.最后一块石头的重量 II

// 有一堆石头，用整数数组 stones 表示。其中 stones[i] 表示第 i 块石头的重量。
// 每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：
// 如果 x == y，那么两块石头都会被完全粉碎；
// 如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
// 最后，最多只会剩下一块 石头。返回此石头 最小的可能重量 。如果没有石头剩下，就返回 0。

// 方法1： 动态规划 + 一维数组
//  分成2组，求差值最小的情况
func lastStoneWeightII(stones []int) int {
	n := len(stones)
	if n == 1 {
		return stones[0]
	}
	// 分成2组差值最小的数
	sum := 0
	for i := 0; i < n; i++ {
		sum += stones[i]
	}
	minRes := int((^uint(0)) >> 1)
	// 搜到一半就行了，超过一般效果一样
	m := sum / 2
	// dp[i][j] 表示前i个数中是否存在和等于j的
	dp := make([]bool, m+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		v := stones[i-1]
		for j := m; j >= 1; j-- {
			if j >= v {
				dp[j] = dp[j] || dp[j-v]
			}
			if dp[j] {
				minRes = min(minRes, sum-2*j)
			}
		}
	}
	return minRes
}

// 方法2： 动态规划 + 二维数组
//  分成2组，求差值最小的情况
func lastStoneWeightII_2(stones []int) int {
	n := len(stones)
	if n == 1 {
		return stones[0]
	}
	// 分成2组差值最小的数
	sum := 0
	for i := 0; i < n; i++ {
		sum += stones[i]
	}
	minRes := int((^uint(0)) >> 1)
	m := sum
	// dp[i][j] 表示前i个数中是否存在和等于j的
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true
	}
	for i := 1; i <= n; i++ {
		v := stones[i-1]
		for j := 1; j <= m; j++ {
			if j >= v {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				dp[i][j] = dp[i-1][j]
			}
			if dp[i][j] {
				tmp := sum - 2*j
				if tmp < 0 {
					tmp = 2*j - sum
				}
				minRes = min(minRes, tmp)
			}
		}
	}
	return minRes
}

// 方法3： 回溯
func solve(stones []int) int {
	n := len(stones)
	if n == 0 {
		return 10000
	}
	if n == 1 {
		return stones[0]
	}
	res := 10000
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			newstones := getNewStone(stones, i, j)
			diff := stones[i] - stones[j]
			if diff < 0 {
				diff = stones[j] - stones[i]
			}
			var tmpRes int
			if diff != 0 {
				newstones = append(newstones, diff)
			}
			tmpRes = solve(newstones)
			res = min(res, tmpRes)
		}
	}
	return res
}

func getNewStone(stones []int, i, j int) []int {
	if len(stones) <= 2 {
		return make([]int, 0)
	}
	newstones := make([]int, len(stones)-2)
	k1, k2 := 0, 0
	for k1 < len(newstones) {
		if k2 == i || k2 == j {
			k2++
			continue
		}
		newstones[k1] = stones[k2]
		k1++
		k2++
	}
	return newstones
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
