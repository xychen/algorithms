package problem11

// 剑指 Offer 11. 旋转数组的最小数字
// 把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。

func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	if numbers[left] < numbers[right] {
		return numbers[left]
	}
	for left < right {
		if numbers[left] == numbers[right] {
			// 启动顺序查找
			minVal := numbers[left]
			for left < right {
				minVal = min(minVal, numbers[left])
				left++
			}
			return minVal
		}
		mid := left + (right-left)/2
		if numbers[mid] >= numbers[left] {
			left = mid
		} else {
			right = mid
		}
		if right-left == 1 {
			return numbers[right]
		}
	}

	return numbers[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
