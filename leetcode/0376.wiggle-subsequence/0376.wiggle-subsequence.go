package problem0376

// 376. 摆动序列
// 如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为 摆动序列 。第一个差（如果存在的话）可能是正数或负数。仅有一个元素或者含两个不等元素的序列也视作摆动序列。
// 给你一个整数数组 nums ，返回 nums 中作为 摆动序列 的 最长子序列的长度 。

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	preDiff, curDiff, ans := 0, 0, 1
	for i := 0; i < n-1; i++ {
		curDiff = nums[i+1] - nums[i]
		if (curDiff > 0 && preDiff <= 0) || (curDiff < 0 && preDiff >= 0) {
			ans++
			preDiff = curDiff
		}
	}
	return ans
}
