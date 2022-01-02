package problem0093

import (
	"strconv"
	"strings"
)

// 93. 复原 IP 地址
// 给定一个只包含数字的字符串，用以表示一个 IP 地址，返回所有可能从 s 获得的 有效 IP 地址 。你可以按任何顺序返回答案。
// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。

func restoreIpAddresses(s string) []string {
	if len(s) < 4 {
		return []string{}
	}
	res := make([]string, 0)
	var backtrack func(s string, startIndex int, path []string)
	backtrack = func(s string, startIndex int, path []string) {
		// if startIndex >= len(s) {
		//     if len(path) == 4 {
		//         res = append(res, strings.Join(path, "."))
		//     }
		//     return
		// }

		// 剪枝
		if len(path) == 3 {
			if startIndex < len(s) && isValicIPElement(s[startIndex:]) {
				path = append(path, s[startIndex:])
				res = append(res, strings.Join(path, "."))
			}
			return
		}
		for i := startIndex; i < len(s); i++ {
			if !isValicIPElement(s[startIndex : i+1]) {
				continue
			}
			path = append(path, s[startIndex:i+1])
			backtrack(s, i+1, path)
			path = path[:len(path)-1]
		}
	}
	path := make([]string, 0)
	backtrack(s, 0, path)
	return res
}

func isValicIPElement(s string) bool {
	if s[0] == '0' && len(s) > 1 {
		return false
	}
	n, err := strconv.Atoi(string(s))
	if err != nil || n > 255 {
		return false
	}
	return true
}
