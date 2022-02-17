package problem40

// 最小的k个数

// 方法一：大顶堆
func getLeastNumbers(arr []int, k int) []int {
	if len(arr) <= k {
		return arr
	}
	if k == 0 {
		return []int{}
	}
	res := make([]int, k)
	copy(res, arr)
	buildHeap(res)
	// fmt.Println(res)

	for i := k; i < len(arr); i++ {
		if arr[i] < res[0] {
			res[0] = arr[i]
			ajustHeap(res, 0)
		}
	}
	return res
}

func buildHeap(arr []int) {
	// 构建大顶堆
	n := len(arr)
	for i := (n - 1) / 2; i >= 0; i-- {
		ajustHeap(arr, i)
	}
}

func ajustHeap(arr []int, start int) {
	// 调整堆
	i, n := start, len(arr)
	for i < n {
		maxChildIndex := 2*i + 1
		if 2*(i+1) < n && arr[2*(i+1)] > arr[maxChildIndex] {
			maxChildIndex = 2 * (i + 1)
		}
		if maxChildIndex < n && arr[i] < arr[maxChildIndex] {
			arr[i], arr[maxChildIndex] = arr[maxChildIndex], arr[i]
			i = maxChildIndex
		} else {
			break
		}
	}
}

// 方法二：快排partition的方式
func getLeastNumbers2(arr []int, k int) []int {
	if len(arr) <= k {
		return arr
	}
	left, right := 0, len(arr)-1
	for {
		index := partition(arr, left, right)
		if index == k {
			return arr[:k]
		} else if index > k {
			right = index - 1
		} else if index < k {
			left = index + 1
		}
	}
	return []int{}
}

func partition(arr []int, left, right int) int {
	i, j := left, right
	v := arr[left]
	for i < j {
		for i < j && arr[j] >= v {
			j--
		}
		for i < j && arr[i] <= v {
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i], arr[left] = arr[left], arr[i]
	return i
}
