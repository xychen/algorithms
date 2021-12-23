package problem53_1

// 剑指 Offer 53 - I. 在排序数组中查找数字 I

func search(nums []int, target int) int {
	if len(nums) <= 0 {
		return 0
	}
	left := leftBound(nums, target)
	if left < 0 || left >= len(nums) {
		return 0
	}
	right := rightBound(nums, target)
	return right - left + 1
}

func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	if left < len(nums) && nums[left] == target {
		return left
	}
	return -1
}

func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	if right >= 0 && nums[right] == target {
		return right
	}
	return -1
}
