package problem0438

//在s中查找所有p的异位词（全排列）

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	need := make(map[byte]int)
	window := make(map[byte]int)

	for i, _ := range p {
		c := p[i]
		need[c]++
	}
	left, right := 0, 0
	valid := 0

	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for right-left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
			}
			c := s[left]
			left++
			if _, ok := need[c]; ok {
				if window[c] == need[c] {
					valid--
				}
				window[c]--
			}
		}
	}
	return res
}
