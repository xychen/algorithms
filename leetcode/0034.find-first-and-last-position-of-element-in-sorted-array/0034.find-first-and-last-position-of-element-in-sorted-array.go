package problem0034

//查找有序数组中等于target的第一个元素和最后一个元素的下标

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	res := []int{-1, -1}
	left, right := 0, len(nums)-1
	//找左边界(终止条件，left = right+1)
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if left >= len(nums) || nums[left] != target {
		return res
	}
	res[0] = left

	left, right = 0, len(nums)-1
	//找右边界 (终止条件 left = right + 1)
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if right < 0 || nums[right] != target {
		return res
	}
	res[1] = right

	return res
}
