package problem0541

// 541. 反转字符串 II

import "strings"

func reverseStr(s string, k int) string {
	slist := strings.Split(s, "")
	left, right, start := 0, 0, 0
	for right < len(s) {
		// 左闭右开
		left, right = start, start+2*k
		if right > len(s) {
			right = len(s)
			if right-left < k {
				reverse(slist, left, right)
			} else if right-left >= k {
				reverse(slist, left, left+k)
			}
		} else {
			reverse(slist, left, left+k)
		}
		start = right
	}
	return strings.Join(slist, "")
}

func reverse(slist []string, start, end int) {
	// 左闭右开
	end -= 1
	for start <= end {
		slist[start], slist[end] = slist[end], slist[start]
		start++
		end--
	}
}
