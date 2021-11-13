package problem0209

// 209. 长度最小的子数组
// 给定一个含有 n 个正整数的数组和一个正整数 target 。
// 找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

func minSubArrayLen(target int, nums []int) int {
	left, right, sum, res := 0, 0, 0, len(nums)+1

	for right < len(nums) {
		sum += nums[right]
		for left <= right && sum >= target {
			res = min(res, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
		if res == 1 {
			break
		}
	}
	if res > len(nums) {
		return 0
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
