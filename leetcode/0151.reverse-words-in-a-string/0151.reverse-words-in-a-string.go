package problem0151

// 151. 翻转字符串里的单词

import "strings"

func reverseWords(s string) string {
	slist := strings.Split(s, " ")
	slist = filterSpace(slist)
	reverse(slist)
	return strings.Join(slist, " ")
}

func reverse(slist []string) {
	left, right := 0, len(slist)-1
	for left <= right {
		slist[left], slist[right] = slist[right], slist[left]
		left++
		right--
	}
}

func filterSpace(slist []string) []string {
	slow, fast := 0, 0
	for fast < len(slist) {
		if slist[fast] == "" {
			fast++
			continue
		}
		slist[slow] = slist[fast]
		slow++
		fast++
	}
	return slist[:slow]
}
