package problem0015

// 15. 三数之和

// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。

//  *** 可参考18题的通用解决方案 ***

import "sort"

func threeSum(nums []int) [][]int {
	// 排序
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		tmpres := twoSum(nums, i+1, -num)
		for _, item := range tmpres {
			item = append(item, num)
			res = append(res, item)
		}
		//避免重复
		for i < len(nums)-1 && nums[i+1] == num {
			i++
		}
	}
	return res
}

func twoSum(nums []int, start, target int) [][]int {
	res := make([][]int, 0)
	l, r := start, len(nums)-1
	for l < r {
		lval, rval := nums[l], nums[r]
		if lval+rval == target {
			tmp := []int{lval, rval}
			res = append(res, tmp)
			for l < r && rval == nums[r] {
				r--
			}
			for l < r && lval == nums[l] {
				l++
			}
		} else if lval+rval > target {
			for l < r && rval == nums[r] {
				r--
			}
		} else if lval+rval < target {
			for l < r && lval == nums[l] {
				l++
			}
		}
	}
	return res
}
