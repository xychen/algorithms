package problem49

// 方法一：使用动态规划
func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p1, p2, p3 := 1, 1, 1

	for i := 2; i <= n; i++ {
		dp[i] = min([]int{dp[p1] * 2, dp[p2] * 3, dp[p3] * 5})
		if dp[i] == dp[p1]*2 {
			p1++
		}
		if dp[i] == dp[p2]*3 {
			p2++
		}
		if dp[i] == dp[p3]*5 {
			p3++
		}
	}
	return dp[n]
}

// 方法二
func nthUglyNumber2(n int) int {
	uglyNumList := []int{1, 2, 3, 4, 5}
	if n <= len(uglyNumList) {
		return uglyNumList[n-1]
	}
	i := len(uglyNumList)
	tmpSel := make([]int, 3)
	for i < n {
		curLen := len(uglyNumList)
		curMax := uglyNumList[curLen-1]

		for j := 1; j < curLen; j++ {
			tmpSel[0] = 2 * uglyNumList[j]
			if tmpSel[0] > curMax {
				break
			}
		}
		for j := 1; j < curLen; j++ {
			tmpSel[1] = 3 * uglyNumList[j]
			if tmpSel[1] > curMax {
				break
			}
		}
		for j := 1; j < curLen; j++ {
			tmpSel[2] = 5 * uglyNumList[j]
			if tmpSel[2] > curMax {
				break
			}
		}
		uglyNumList = append(uglyNumList, min(tmpSel))
		i++
	}
	return uglyNumList[len(uglyNumList)-1]
}

func min(nums []int) int {
	res := nums[0]
	for _, num := range nums {
		if num < res {
			res = num
		}
	}
	return res
}
