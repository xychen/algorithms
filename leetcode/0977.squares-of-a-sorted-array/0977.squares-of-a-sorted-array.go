package problem0977

// 977. 有序数组的平方
// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

func sortedSquares(nums []int) []int {
	// 负数结束的位置
	flag := -1
	for i, num := range nums {
		if num < 0 {
			flag = i
		}
		nums[i] = num * num
	}
	res := make([]int, len(nums))
	left := flag
	right := flag + 1
	i := 0
	for left >= 0 || right < len(nums) {
		if left >= 0 && right < len(nums) {
			if nums[left] < nums[right] {
				res[i] = nums[left]
				left--
			} else {
				res[i] = nums[right]
				right++
			}
			i++
			continue
		}
		if left >= 0 {
			res[i] = nums[left]
			left--
		}
		if right < len(nums) {
			res[i] = nums[right]
			right++
		}
		i++
	}
	return res
}
