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

//相关题目：
//1. 判断1个整数是不是2的整数次方   解： 2的整数次方二进制表示中只有一个1，和它-1后的数相与，等于0则是
//2. 输入2个数m和n，计算需要改变m的二进制表示中多少位才能得到n   解： m和n相异或，再求异或后二进制中1的数量
