package problem004

import "testing"

type param struct {
	one string
}

type ans struct {
	res string
}

type question struct {
	p param
	a ans
}

func Test_OK(t *testing.T) {
	qs := []question{
		question{p: param{one: "we are happy"}, a: ans{res: "we%20are%20happy"}},
	}

	for _, q := range qs {
		if q.a.res != ReplaceBlank(q.p.one) {
			t.Errorf("输入%v", q)
		}
	}
}
