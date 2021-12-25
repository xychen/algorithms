package problem15

// 剑指 Offer 15. 二进制中1的个数

func hammingWeight(num uint32) int {
	c := 0
	for num > 0 {
		num = num & (num - 1)
		c++
	}
	return c
}
