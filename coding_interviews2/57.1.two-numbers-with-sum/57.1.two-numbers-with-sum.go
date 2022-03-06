package problem57_1

// 剑指 Offer 57. 和为s的两个数字
// 输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		tmpSum := nums[left] + nums[right]
		if tmpSum == target {
			return []int{nums[left], nums[right]}
		} else if tmpSum > target {
			right--
		} else if tmpSum < target {
			left++
		}
	}
	return []int{}
}

// func twoSum(nums []int, target int) []int {

//     for i, num := range nums {
//         j := search(nums, target-num, i+1, len(nums)-1)
//         if j == -1 {
//             continue
//         }
//         return []int{num, nums[j]}
//     }
//     return []int{}
// }

// // return index
// func search(nums []int, target, left, right int) int {
//     for left <= right {
//         mid := left + (right-left)/2
//         if nums[mid] == target {
//             return mid
//         } else if nums[mid] > target {
//             right = mid-1
//         } else if nums[mid] < target {
//             left = mid+1
//         }
//     }

//     return -1
// }
