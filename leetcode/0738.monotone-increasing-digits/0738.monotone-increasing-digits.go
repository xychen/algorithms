package problem0738

// 738. 单调递增的数字
// 给定一个非负整数 N，找出小于或等于 N 的最大的整数，同时这个整数需要满足其各个位数上的数字是单调递增。

// （当且仅当每个相邻位数上的数字 x 和 y 满足 x <= y 时，我们称这个整数是单调递增的。）

func monotoneIncreasingDigits(n int) int {
	nums := make([]uint8, 0)
	// 数字转数组
	for n > 0 {
		nums = append(nums, uint8(n%10))
		n = n / 10
	}
	flag := -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			nums[i+1]--
			// flag以前的数字都有转成9
			flag = i
		}
	}
	ans := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if i <= flag {
			nums[i] = uint8(9)
		}
		ans = ans*10 + int(nums[i])
	}
	return ans
}
