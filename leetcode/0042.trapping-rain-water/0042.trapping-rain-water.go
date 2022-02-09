package problem0042

// 42. 接雨水
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

// 空间复杂度为O(1)的解法
func trap(height []int) int {
	n := len(height)
	left, right := 0, n-1
	lmax, rmax := height[0], height[n-1]
	res := 0
	for left < right {
		lmax = max(lmax, height[left])
		rmax = max(rmax, height[right])
		if lmax < rmax {
			res += lmax - height[left]
			left++
		} else {
			res += rmax - height[right]
			right--
		}
	}
	return res
}

// 空间复杂度为O(n)的解法，提前计算leftMax和rightMax
func trap2(height []int) int {
	n := len(height)
	lmax := make([]int, n)
	rmax := make([]int, n)
	lmax[0] = height[0]
	for i := 1; i < n; i++ {
		lmax[i] = max(lmax[i-1], height[i])
	}
	rmax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rmax[i] = max(rmax[i+1], height[i])
	}
	ans := 0
	for i := 1; i < n-1; i++ {
		ans += min(lmax[i], rmax[i]) - height[i]
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
