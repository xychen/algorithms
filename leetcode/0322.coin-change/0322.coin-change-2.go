package problem0322

// 零钱兑换
// 方法二： 回溯 + 记忆化搜索
const INT_MAX = int((^uint(0) >> 1))

var MAX_RES int
var dp map[int]int

func coinChange2(coins []int, amount int) int {
	MAX_RES = amount + 1
	dp = make(map[int]int)
	res := solve(coins, amount)
	if res == MAX_RES {
		return -1
	}
	return res
}

func solve(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return MAX_RES
	}
	if v, ok := dp[amount]; ok {
		return v
	}
	minNum := MAX_RES
	for _, coin := range coins {
		//最选择
		res := solve(coins, amount-coin)
		minNum = min(minNum, res+1)
		//回退
		//…… do noting
	}
	dp[amount] = minNum
	return minNum
}
