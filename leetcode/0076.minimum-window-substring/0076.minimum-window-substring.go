package problem0076

import (
	"strings"
)

const MAX_INT = 1000000000

func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	sarr := strings.Split(s, "")
	tarr := strings.Split(t, "")
	//初始化need数组
	need := make(map[string]int)
	window := make(map[string]int)
	for _, c := range tarr {
		// if _, ok := need[c]; ok {
		// 	need[c]++
		// } else {
		// 	need[c] = 1
		// }
		need[c]++
		window[c] = 0
	}

	valid := 0
	start := 0
	length := MAX_INT
	left, right := 0, 0
	for right < len(sarr) {
		tmps := sarr[right]
		right++
		//需要
		if _, ok := need[tmps]; ok {
			window[tmps]++
			if need[tmps] == window[tmps] {
				valid++
			}
		}
		//数量还不够
		if valid < len(need) {
			continue
		}
		//数量够了,试图缩减
		for valid >= len(need) {
			//更新数据
			if right-left < length {
				start = left
				length = right - left
			}
			tmps = sarr[left]
			left++
			if _, ok := need[tmps]; ok {
				if need[tmps] == window[tmps] {
					valid--
				}
				window[tmps]--
			}
		}

	}
	if length == MAX_INT {
		return ""
	}
	return strings.Join(sarr[start:start+length], "")
}
