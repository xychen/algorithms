package problem53_2

// 剑指 Offer 53 - II. 0～n-1中缺失的数字
func missingNumber(nums []int) int {
	if len(nums) < 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if mid == nums[mid] {
			left = mid + 1
		} else if nums[mid] > mid {
			right = mid - 1
		}
	}
	return left
}
