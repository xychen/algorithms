package problem0027

// 27. 移除元素

func removeElement(nums []int, val int) int {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] == val {
			right++
		} else {
			nums[left] = nums[right]
			left++
			right++
		}
	}
	return left
}
