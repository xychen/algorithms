package problem0031

// 31. 下一个排列

func nextPermutation(nums []int) {
	if len(nums) <= 0 {
		return
	}
	// 从后往前找到第一个较小的元素,【注意是 <=】
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	j := len(nums) - 1
	// 找到了这样的元素
	if i >= 0 {
		// 找到第一个比nums[i]大的数，【注意是 <=】
		for j > i && nums[j] <= nums[i] {
			j--
		}
		// 交换i、j的位置
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 将i+1之后的数据翻转成升序
	left, right := i+1, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
