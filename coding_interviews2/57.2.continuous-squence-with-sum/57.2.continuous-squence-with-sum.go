package problem57_2

// 剑指 Offer 57 - II. 和为s的连续正数序列
// 输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

// 序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

func findContinuousSequence(target int) [][]int {
	if target == 1 {
		return [][]int{}
	}
	left, right := 1, 2
	sum := 3
	res := make([][]int, 0)
	for right < target-1 {
		if sum == target {
			res = append(res, constructSlice(left, right))
			sum -= left
			left++
		} else if sum < target {
			right++
			sum += right
		} else if sum > target {
			sum -= left
			left++
		}
	}
	return res
}

func constructSlice(left, right int) []int {
	n := right - left + 1
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = left
		left++
	}
	return res
}
