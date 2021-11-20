package problem0349

// 349. 两个数组的交集

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums2) < len(nums1) {
		nums1, nums2 = nums2, nums1
	}
	ht := make(map[int]bool)
	for _, num := range nums1 {
		ht[num] = false
	}
	res := make([]int, 0)
	for _, num := range nums2 {
		flag, ok := ht[num]
		if !ok {
			continue
		}
		// 未添加到过结果集
		if !flag {
			res = append(res, num)
			ht[num] = true
		}
	}
	return res
}
