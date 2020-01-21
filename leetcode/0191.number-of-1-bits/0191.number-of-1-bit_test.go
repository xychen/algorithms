package problem0191

import "testing"

type param struct {
	n uint32
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
		res := hammingWeight(q.p.n)
		if res != q.a.res {
			t.Errorf("hammingWeight输入：%v", q)
		}
	}
}

func Test_NumberOf1_2(t *testing.T) {
	for _, q := range qs {
		res := hammingWeight2(q.p.n)
		if res != q.a.res {
			t.Errorf("hammingWeight2输入：%v", q)
		}
	}
}
