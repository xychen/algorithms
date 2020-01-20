package problem0003

import "testing"

type param struct {
	one string
}
type ans struct {
	res int
}

type question struct {
	p param
	a ans
}

var qs = []question{
	question{p: param{one: "pwwkew"}, a: ans{res: 3}},
	question{p: param{one: "abcabcbb"}, a: ans{res: 3}},
	question{p: param{one: "bbbbb"}, a: ans{res: 1}},
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	for _, q := range qs {
		if q.a.res != lengthOfLongestSubstring(q.p.one) {
			t.Errorf("lengthOfLongestSubstring: %s", q.p.one)
		}
	}
}

func Test_lengthOfLongestSubstring2(t *testing.T) {
	for _, q := range qs {
		if q.a.res != lengthOfLongestSubstring2(q.p.one) {
			t.Errorf("lengthOfLongestSubstring: %s", q.p.one)
		}
	}
}
