package problem0015

// 15. 三数之和

// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。

//  *** 可参考18题的通用解决方案 ***

import "sort"

func threeSum(nums []int) [][]int {
	// 排序
	sort.Ints(nums)
	return nSum(nums, 3, 0, 0)
}

// n表示nsum
func nSum(nums []int, n, start, target int) [][]int {
	res := make([][]int, 0)
	// twosum 是base case
	if n == 2 {
		l, r := start, len(nums)-1
		for l < r {
			lval, rval := nums[l], nums[r]
			sum := lval + rval
			if target < sum {
				for l < r && nums[r] == rval {
					r--
				}
			} else if target > sum {
				for l < r && nums[l] == lval {
					l++
				}
			} else {
				res = append(res, []int{lval, rval})
				for l < r && nums[r] == rval {
					r--
				}
				for l < r && nums[l] == lval {
					l++
				}
			}
		}
		return res
	}

	for i := start; i < len(nums); i++ {
		num := nums[i]
		tmpres := nSum(nums, n-1, i+1, target-num)
		for _, item := range tmpres {
			item = append(item, num)
			res = append(res, item)
		}
		for i < len(nums)-1 && num == nums[i+1] {
			i++
		}
	}
	return res
}
