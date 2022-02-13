package problem0454

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
