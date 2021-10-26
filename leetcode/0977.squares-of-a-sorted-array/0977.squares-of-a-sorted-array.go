package problem0977

// 977. 有序数组的平方
// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

func sortedSquares(nums []int) []int {
	left, right, index := 0, len(nums)-1, len(nums)-1
	res := make([]int, len(nums))

	for left <= right {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			res[index] = nums[left] * nums[left]
			left++
		} else {
			res[index] = nums[right] * nums[right]
			right--
		}
		index--
	}
	return res
}
