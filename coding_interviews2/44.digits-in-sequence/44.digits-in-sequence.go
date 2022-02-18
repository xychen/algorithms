package problem44

// 剑指 Offer 44. 数字序列中某一位的数字
// 数字以0123456789101112131415…的格式序列化到一个字符序列中。在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。

import "math"

func findNthDigit(n int) int {
	if n < 0 {
		return -1
	}
	digits := 1
	for {
		numbers := countOfIntegers(digits)
		if n < numbers*digits {
			return findAtN(n, digits)
		}
		n -= numbers * digits
		digits++
	}
	return -1
}

func countOfIntegers(digits int) int {
	if digits == 1 {
		return 10
	}
	count := int(math.Pow(10, float64(digits-1)))
	return 9 * count
}

func findAtN(n, digits int) int {
	number := beginNumber(digits) + n/digits
	indexFromRight := digits - n%digits
	for i := 1; i < indexFromRight; i++ {
		number /= 10
	}
	return number % 10
}

func beginNumber(digits int) int {
	if digits == 1 {
		return 0
	}
	return int(math.Pow(10, float64(digits-1)))
}
