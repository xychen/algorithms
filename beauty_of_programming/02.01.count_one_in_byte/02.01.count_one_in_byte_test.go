package problem02_01

import (
	"testing"
)

type param struct {
	n byte
}
type ans struct {
	res int
}

type question struct {
	p param
	a ans
}

func Test_OK(t *testing.T) {
	qs := []question{
		question{p: param{n: 0}, a: ans{res: 0}},
		question{p: param{n: 2}, a: ans{res: 1}},
		question{p: param{n: 7}, a: ans{res: 3}},
		question{p: param{n: 15}, a: ans{res: 4}},
		question{p: param{n: 255}, a: ans{res: 8}},
	}

	for _, q := range qs {
		if q.a.res != Count(q.p.n) {
			t.Errorf("输入：%v", q.p)
		}
	}
}
