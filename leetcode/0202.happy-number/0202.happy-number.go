package problem0202

// 202. 快乐数

func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	ht := make(map[int]bool)
	for {
		res := getVal(n)
		if res == 1 {
			return true
		}
		if _, ok := ht[res]; ok {
			break
		}
		ht[res] = true
		n = res
	}
	return false
}

func getVal(n int) int {
	sum := 0
	for n > 0 {
		t := n % 10
		sum += t * t
		n = n / 10
	}
	return sum
}
