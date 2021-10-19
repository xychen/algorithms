package problem0560

// 560. 和为 K 的子数组

// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回该数组中和为 k 的连续子数组的个数。

func subarraySum(nums []int, k int) int {
	presum := make(map[int]int)
	presum[0] = 1
	sumi, sumj := 0, 0
	ans := 0
	for i := 0; i < len(nums); i++ {
		sumi += nums[i]
		sumj = sumi - k
		if v, ok := presum[sumj]; ok {
			ans += v
		}
		presum[sumi] += 1
	}
	return ans
}
