package problem0001

// 1. 两数之和

func twoSum(nums []int, target int) []int {
	ht := make(map[int]int)
	for index, num := range nums {
		k := target - num
		if v, ok := ht[k]; ok {
			return []int{v, index}
		}
		ht[num] = index
	}
	return []int{-1, -1}
}
