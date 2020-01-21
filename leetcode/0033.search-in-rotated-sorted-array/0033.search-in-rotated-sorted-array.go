package problem0033

//直接二分查找
func Search1(nums []int, target int) int {
	n := len(nums)
	if n <= 0 {
		return -1
	}
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[left] { //mid在左侧，左侧是递增
			if nums[left] <= target && target < nums[mid] { //在左侧找
				right = mid - 1
			} else { //只能从右边找
				left = mid + 1
			}
		} else { //mid在右侧，右侧是递增
			if nums[mid] < target && target <= nums[right] { //在右侧找
				left = mid + 1
			} else { //只能从左侧找
				right = mid - 1
			}
		}
	}

	return -1
}

//将mid转成原数组的index
func Search2(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}
	n := len(nums)
	offset := minIndexInRotateArray(nums)
	left, right := offset, n-1+offset

	for left <= right {
		mid := (left + right) / 2
		realMid := mid % n
		if nums[realMid] == target {
			return realMid
		} else if nums[realMid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func Search3(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}
	n := len(nums)
	minIndex := minIndexInRotateArray(nums)
	var left, right int
	if target > nums[n-1] { //左边找
		left, right = 0, minIndex-1
	} else { //右边找
		left, right = minIndex, n-1
	}

	//二分查找
	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

//查找拐点的下标
func minIndexInRotateArray(nums []int) int {
	if len(nums) <= 0 {
		return -1
	}
	left, mid, right := 0, 0, len(nums)-1
	for nums[left] > nums[right] {
		if right-left == 1 {
			mid = right
			break
		}
		mid = (left + right) / 2
		// left 或 right 等于mid的时候，只有left=right或left与right相邻的时候才会出现，但上边已经确保不会出去，
		// 所以赋值时不需要对mid±1
		if nums[mid] >= nums[left] {
			left = mid
		} else if nums[mid] <= nums[right] {
			right = mid
		}
	}
	return mid
}
