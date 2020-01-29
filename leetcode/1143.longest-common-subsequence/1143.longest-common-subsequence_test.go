package problem1143

import "testing"

type param struct {
	text1 string
	text2 string
}

type ans struct {
	res int
}

type question struct {
	p param
	a ans
}

var qs = []question{
	question{p: param{text1: "abcde", text2: "ace"}, a: ans{res: 3}},
	question{p: param{text1: "abc", text2: "abc"}, a: ans{res: 3}},
	question{p: param{text1: "abc", text2: "def"}, a: ans{res: 0}},
	question{p: param{text1: "", text2: ""}, a: ans{res: 0}},
	question{p: param{text1: "", text2: "abc"}, a: ans{res: 0}},
}

func Test_longestCommonSubsequence(t *testing.T) {
	for i, q := range qs {
		if longestCommonSubsequence(q.p.text1, q.p.text2) != q.a.res {
			t.Errorf("longestCommonSubsequence 第 %d 个输入报错", i)
		}
	}
}

func Test_longestCommonSubsequence_2(t *testing.T) {
	for i, q := range qs {
		if longestCommonSubsequence(q.p.text1, q.p.text2) != q.a.res {
			t.Errorf("longestCommonSubsequence_2 第 %d 个输入报错", i)
		}
	}
}
