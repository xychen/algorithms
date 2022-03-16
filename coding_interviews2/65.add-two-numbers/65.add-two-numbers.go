package problem65

// 剑指 Offer 65. 不用加减乘除做加法

func add(a int, b int) int {
	sum, carry := 0, 0
	for {
		sum = a ^ b
		carry = (a & b) << 1
		a = sum
		b = carry
		if carry == 0 {
			break
		}
	}
	return sum
}
