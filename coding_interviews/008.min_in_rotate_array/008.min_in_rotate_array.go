package problem008

//旋转数组中的最小值	ps:只考虑严格递增的场景
func MinInRotateArr(arr []int) int {
	low, mid, high := 0, 0, len(arr)-1
	for arr[low] > arr[high] {
		if high-low == 1 {
			mid = high
			break
		}
		mid = (low + high) / 2
		if arr[mid] >= arr[low] {
			low = mid
		} else if arr[mid] <= arr[high] {
			high = mid
		}
	}
	return arr[mid]
}
