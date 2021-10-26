package problem1005

// 1005. K 次取反后最大化的数组和
// 给定一个整数数组 A，我们只能用以下方法修改该数组：我们选择某个索引 i 并将 A[i] 替换为 -A[i]，然后总共重复这个过程 K 次。（我们可以多次选择同一个索引 i。）

// 以这种方式修改数组后，返回数组可能的最大和。

import (
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	// 按绝对值从大到小排序
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})
	n := len(nums)
	for i := 0; i < n && k > 0; i++ {
		if nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}
	}
	if k%2 == 1 {
		nums[n-1] = -nums[n-1]
	}
	return sum(nums)
}

func sum(nums []int) int {
	ans := 0
	for _, n := range nums {
		ans += n
	}
	return ans
}
