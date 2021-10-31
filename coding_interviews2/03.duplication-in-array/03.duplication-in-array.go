package problem03

// 剑指 Offer 03. 数组中重复的数字
// 找出数组中重复的数字。

// 在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for {
			if nums[i] == i {
				break
			}
			// 存在重复
			if nums[nums[i]] == nums[i] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}
