package problem46

// 剑指 Offer 46. 把数字翻译成字符串
// 给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

func translateNum(num int) int {
	if num == 0 {
		return 1
	} else if num < 0 {
		return 0
	}
	nums := covertToArray(num)
	// dp := make([]int, len(nums))
	// dp[0] = 1
	prepre := 0
	pre := 1
	res := 1
	for i := 1; i < len(nums); i++ {
		tmp := nums[i-1]*10 + nums[i]
		res = pre
		if tmp <= 25 && tmp >= 10 {
			if i-2 >= 0 {
				res += prepre
			} else {
				// i=1的时候的计算方法
				res += 1
			}
		}
		prepre = pre
		pre = res
	}
	return res
}

func covertToArray(num int) []int {
	res := make([]int, 0)
	for num > 0 {
		res = append(res, num%10)
		num = num / 10
	}
	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}
