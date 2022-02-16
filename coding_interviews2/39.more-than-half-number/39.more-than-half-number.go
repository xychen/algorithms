package problem39

// 方法一：
// 遇到相同的数，次数+1， 遇到不同的数，次数-1，当次数减到0的时候，需要换一个数字
// 因为要找的数字比其他数字总和还要多，所以最后能剩下
func majorityElement(nums []int) int {
	res := nums[0]
	times := 1
	for i := 1; i < len(nums); i++ {
		if res == nums[i] {
			times++
		} else {
			times--
		}
		if times == 0 {
			res = nums[i]
			times = 1
		}
	}
	return res
}

// 方法二：partition方式
func majorityElement2(nums []int) int {
	left, right := 0, len(nums)-1
	mid := len(nums) / 2
	for left <= right {
		index := partition(nums, left, right)
		if mid == index {
			return nums[mid]
		} else if mid > index {
			left = index + 1
		} else if mid < index {
			right = index - 1
		}
	}
	return -1
}

func partition(nums []int, left, right int) int {
	v := nums[left]
	i, j := left, right
	// 等于v的放到右边
	for i < j {
		for i < j && nums[j] >= v {
			j--
		}
		for i < j && nums[i] <= v {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[i], nums[left] = nums[left], nums[i]
	return i
}
