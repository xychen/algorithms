package problem0454

// 454. 四数相加 II
// 给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：

// 0 <= i, j, k, l < n
// nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	ht := make(map[int]int, 0)
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			ht[n1+n2] += 1
		}
	}
	ans := 0
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			k := 0 - (n3 + n4)
			if v, ok := ht[k]; ok {
				ans += v
			}
		}
	}
	return ans
}
