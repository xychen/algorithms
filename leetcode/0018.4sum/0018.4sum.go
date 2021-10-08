package problem0018

// 18. 四数之和
// 给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] ：

// 0 <= a, b, c, d < n
// a、b、c 和 d 互不相同
// nums[a] + nums[b] + nums[c] + nums[d] == target
// 你可以按 任意顺序 返回答案 。
import "sort"

func fourSum(nums []int, target int) [][]int {
	// 第一步要排序
	sort.Ints(nums)
	return nSum(nums, 4, 0, target)
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
