package problem010

import "testing"

type param struct {
	n uint
}

type ans struct {
	res int
}

type question struct {
	p param
	a ans
}

var qs = []question{
	question{p: param{n: 0}, a: ans{res: 0}},
	question{p: param{n: 1}, a: ans{res: 1}},
	question{p: param{n: 2}, a: ans{res: 1}},
}

func Test_NumberOf1(t *testing.T) {
	for _, q := range qs {
		res := NumberOf1(q.p.n)
		if res != q.a.res {
			t.Errorf("NumberOf1输入：%v", q)
		}
	}
}

func Test_NumberOf1_2(t *testing.T) {
	for _, q := range qs {
		res := NumberOf1_2(q.p.n)
		if res != q.a.res {
			t.Errorf("NumberOf1_2输入：%v", q)
		}
	}
}
