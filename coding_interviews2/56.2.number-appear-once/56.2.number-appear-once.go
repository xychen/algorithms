package problem56_2

// 剑指 Offer 56 - II. 数组中数字出现的次数 II

func singleNumber(nums []int) int {
	resList := make([]int, 32)

	for _, num := range nums {
		for i := 31; i >= 0; i-- {
			resList[i] = (resList[i] + (num>>i)&1) % 3
		}
	}
	res := 0
	for i := 31; i >= 0; i-- {
		res = res<<1 + resList[i]
	}
	return res
}
