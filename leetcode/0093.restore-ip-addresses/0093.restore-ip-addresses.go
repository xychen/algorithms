package problem0093

// 93. 复原 IP 地址
// 给定一个只包含数字的字符串，用以表示一个 IP 地址，返回所有可能从 s 获得的 有效 IP 地址 。你可以按任何顺序返回答案。
// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。

import (
	"strconv"
	"strings"
)

var res []string

func restoreIpAddresses(s string) []string {
	l := len(s)
	res := make([]string, 0)
	if l < 4 || l > 12 {
		return res
	}

	var solve func(s string, part int, strs []string)
	solve = func(s string, part int, strs []string) {
		if len(s) == 0 {
			return
		}
		if part == 4 {
			strs = append(strs, s)
			if isValidIP(strs) {
				res = append(res, strings.Join(strs, "."))
			}
			return
		}

		for i := 0; i < len(s); i++ {
			//每一段最多3个数字，剪枝
			if i > 3 {
				break
			}
			tmpstr := s[:i+1]
			strs = append(strs, tmpstr)
			solve(s[i+1:], part+1, strs)
			strs = strs[:len(strs)-1]
		}
	}
	strs := make([]string, 0)
	solve(s, 1, strs)
	return res
}

func isValidIP(strs []string) bool {
	if len(strs) != 4 {
		return false
	}
	for _, str := range strs {
		if !isValidIPSlice(str) {
			return false
		}
	}
	return true
}
func isValidIPSlice(str string) bool {
	// 0开头的只能是0，不能是01这种
	if len(str) != 1 && str[0] == '0' {
		return false
	}
	// 最大是255， 大于3位的不符合
	if len(str) > 3 {
		return false
	}
	v, _ := strconv.Atoi(str)
	if v > 255 {
		return false
	}
	return true
}
