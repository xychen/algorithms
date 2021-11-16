package problem0242

// 242. 有效的字母异位词

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ht := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		ht[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		k := t[i]
		v, ok := ht[k]
		if !ok {
			return false
		}
		if v <= 0 {
			return false
		}
		ht[k]--
	}
	return true
}
