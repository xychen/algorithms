package problem45

// 剑指 Offer 45. 把数组排成最小的数
// 输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

import (
	"sort"
	"strconv"
	"strings"
)

func minNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}
	sort.Slice(strs, func(i, j int) bool {
		a := strs[i] + strs[j]
		b := strs[j] + strs[i]
		if a < b {
			return true
		}
		return false
	})

	return strings.Join(strs, "")
}
