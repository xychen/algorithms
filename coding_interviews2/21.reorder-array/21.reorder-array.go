package problem21

// 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面

func exchange(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		for (left < right) && (nums[left]&0x1) == 1 {
			left++
		}
		for (left < right) && (nums[right]&0x1) == 0 {
			right--
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	return nums
}
