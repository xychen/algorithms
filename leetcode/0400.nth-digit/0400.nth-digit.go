package problem0400

// 400. 第 N 位数字
// 给你一个整数 n ，请你在无限的整数序列 [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...] 中找出并返回第 n 位数字。

import "math"

func findNthDigit(n int) int {
	digit := 1
	for {
		numbers := digit * countNumbersByDigit(digit)
		if numbers > n {
			return doFindNthDigit(digit, n)
		}
		n -= numbers
		digit++
	}
	return -1
}

func doFindNthDigit(digit, n int) int {
	number := beginNumberOfDigit(digit) + n/digit

	for i := 1; i < digit-n%digit; i++ {
		number = number / 10
	}
	return number % 10
}

func beginNumberOfDigit(digit int) int {
	if digit == 1 {
		return 0
	}
	return int(math.Pow(10, float64(digit-1)))
}

func countNumbersByDigit(digit int) int {
	if digit == 1 {
		return 10
	}
	return 9 * int(math.Pow(10, float64(digit-1)))
}
