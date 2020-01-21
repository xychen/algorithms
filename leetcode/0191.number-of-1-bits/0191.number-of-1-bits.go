package problem0191

//与《剑指offer》第10题相同
//与《编程之美》第2.1节相同


//普通解法，不断右移和1进行与运算
func hammingWeight(num uint32) int {
	var count uint32 = 0
	for num > 0 {
		count += num & 0x1
		num >>= 1
	}
	return int(count)
}


//与减1的数进行与操作，相当于把最右边的1变成0
func hammingWeight2(num uint32) int {
	var count uint = 0
	for num > 0 {
		count += 1
		num = num & (num - 1)
	}
	return int(count)
}
