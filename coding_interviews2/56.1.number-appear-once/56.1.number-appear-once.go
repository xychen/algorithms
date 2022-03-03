package problem56_1

func singleNumbers(nums []int) []int {

	if len(nums) <= 0 {
		return []int{}
	}
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		m = m ^ nums[i]
	}
	k := findK(m)
	m1, m2 := 0, 0
	for _, num := range nums {
		if num&k == 0 {
			m1 = m1 ^ num
		} else {
			m2 = m2 ^ num
		}
	}
	return []int{m1, m2}
}

func findK(n int) int {
	k := 1
	for n&k == 0 {
		k = k << 1
	}
	return k
}
