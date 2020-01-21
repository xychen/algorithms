package problem014

//调整数组顺序，使奇数在前，偶数在后
func ReorderArray(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		for left < right && nums[left]&0x1 == 1 {
			left += 1
		}
		for left < right && nums[right]&0x1 == 0 {
			right -= 1
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	return nums
}
