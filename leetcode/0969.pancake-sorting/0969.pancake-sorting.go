package problem0969

// 969. 煎饼排序

var res []int

func pancakeSort(arr []int) []int {
	res = make([]int, 0)
	solve(&arr, len(arr))
	return res
}

func solve(arr *[]int, n int) {
	if n == 1 {
		return
	}
	//最大的放最下边
	maxIndex := findMaxIndex(*arr, n)
	reverse(arr, 0, maxIndex)
	res = append(res, maxIndex+1)
	reverse(arr, 0, n-1)
	res = append(res, n)

	//递归调用，对上边n-1个继续排
	solve(arr, n-1)
}

func reverse(arr *[]int, left, right int) {
	for left < right {
		(*arr)[left], (*arr)[right] = (*arr)[right], (*arr)[left]
		left++
		right--
	}
}

func findMaxIndex(arr []int, n int) int {
	maxIndex, maxVal := 0, arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > maxVal {
			maxIndex = i
			maxVal = arr[i]
		}
	}
	return maxIndex
}
