package problem58_1

import "strings"

// 剑指 Offer 58 - I. 翻转单词顺序

func reverseWords(s string) string {
	if len(s) == 0 {
		return ""
	}
	words := strings.Split(s, " ")
	slow, fast := 0, 0
	for fast < len(words) {
		if words[fast] == "" {
			fast++
			continue
		}
		words[slow] = words[fast]
		slow++
		fast++
	}
	left, right := 0, slow-1
	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}
	words = words[:slow]
	return strings.Join(words, " ")
}
