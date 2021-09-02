package problem0567

import (
	"strings"
)

func checkInclusion(s1 string, s2 string) bool {
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
	// fmt.Println("need:", len(need), need)
	s2arr := strings.Split(s2, "")
	left, right := 0, 0
	window := make(map[string]int)
	valid := 0
	for right < len(s2arr) {
		c := s2arr[right]
		window[c]++
		right++
		if _, ok := need[c]; ok && window[c] == need[c] {
			valid++
		}
		// fmt.Println(len(window), ":", window)
		for valid >= len(need) {
			if checkValid(window, need) {
				return true
			}
			c := s2arr[left]
			left++
			if _, ok := need[c]; ok && window[c] == need[c] {
				valid--
			}
			window[c]--
			if window[c] == 0 {
				delete(window, c)
			}
		}

	}

	return false
}

func checkValid(window, need map[string]int) bool {
	if len(window) != len(need) {
		return false
	}

	for k, v := range window {
		if v != need[k] {
			return false
		}
	}
	return true
}
