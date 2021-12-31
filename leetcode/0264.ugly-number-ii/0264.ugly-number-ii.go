package problem0264

// 264. 丑数 II
// 给你一个整数 n ，请你找出并返回第 n 个 丑数 。

// 丑数 就是只包含质因数 2、3 和/或 5 的正整数。

func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	p1, p2, p3 := 1, 1, 1
	nums := make([]int, n+1)
	nums[1] = 1
	for i := 2; i <= n; i++ {
		nums[i] = min([]int{nums[p1] * 2, nums[p2] * 3, nums[p3] * 5})
		if nums[i] == nums[p1]*2 {
			p1++
		}
		if nums[i] == nums[p2]*3 {
			p2++
		}
		if nums[i] == nums[p3]*5 {
			p3++
		}
	}
	return nums[n]
}

func min(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < res {
			res = nums[i]
		}
	}
	return res
}
