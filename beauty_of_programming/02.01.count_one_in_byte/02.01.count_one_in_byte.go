package problem02_01

//计算byte中1的个数
func Count(n uint8) int {
	var c uint8 = 0
	for n != 0 {
		c += n & 0x1
		n >>= 1
	}
	return int(c)
}
