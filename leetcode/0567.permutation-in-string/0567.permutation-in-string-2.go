package problem0567

import (
	"strings"
)

func checkInclusion2(s1 string, s2 string) bool {
	if s1 == "" {
		return true
	}
	if s2 == "" {
		return false
	}
	s1arr := strings.Split(s1, "")
	need := make(map[string]int)
	for _, c := range s1arr {
		need[c]++
	}
	s2arr := strings.Split(s2, "")
	left, right := 0, 0
	window := make(map[string]int)
	valid := 0

	for right < len(s2arr) {
		c := s2arr[right]
		right++
		//不符合的c就不加入window中了，因为下边的for循环用 right - left判断，如果中间夹了一个不符合的，left会往前走，导致某个符合的字符也不符合了
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for right-left >= len(s1arr) {
			if valid == len(need) {
				return true
			}
			c := s2arr[left]
			left++

			if _, ok := need[c]; ok {
				if window[c] == need[c] {
					valid--
				}
				window[c]--
			}
		}
	}

	return false
}
