package problem0287

// 287. 寻找重复数
// 给定一个包含 n + 1 个整数的数组 nums ，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。

// 假设 nums 只有 一个重复的整数 ，找出 这个重复的数 。

// 你设计的解决方案必须不修改数组 nums 且只用常量级 O(1) 的额外空间。

func findDuplicate(nums []int) int {
	start, end := 1, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		// start~mid的数量
		c := countRange(nums, start, mid)
		if start == end {
			return start
		}
		// 数量超了
		if c > (mid - start + 1) {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return -1
}

func countRange(nums []int, start, end int) int {
	c := 0
	for _, num := range nums {
		if num >= start && num <= end {
			c++
		}
	}
	return c
}
