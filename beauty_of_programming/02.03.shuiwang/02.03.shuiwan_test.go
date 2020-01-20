package problem02_03

import (
	"testing"
)

type param struct {
	a []int
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
		question{p: param{a: []int{1, 2, 2, 2, 3, 3, 2}}, a: ans{2}},
	}

	for _, q := range qs {
		if q.a.res != Find(q.p.a) {
			t.Errorf("输入:%v", q.p)
		}
	}
}
