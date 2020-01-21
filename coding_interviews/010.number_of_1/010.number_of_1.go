package problem010

//普通解法，不断右移和1进行与运算
func NumberOf1(n uint) int {
	var count uint = 0
	for n > 0 {
		count += n & 0x1
		n >>= 1
	}
	return int(count)
}

//与减1的数进行与操作，相当于把最右边的1变成0
func NumberOf1_2(n uint) int {
	var count uint = 0
	for n > 0 {
		count += 1
		n = n & (n - 1)
	}
	return int(count)
}
