package problem009

import "testing"

type param struct {
	n int
}

type ans struct {
	res uint64
}

type question struct {
	p param
	a ans
}

func Test_Fibonacci(t *testing.T) {
	qs := []question{
		question{p: param{n: 0}, a: ans{res: 0}},
		question{p: param{n: 1}, a: ans{res: 1}},
		question{p: param{n: 2}, a: ans{res: 1}},
		question{p: param{n: 3}, a: ans{res: 2}},
	}

	for _, q := range qs {
		if Fibonacci(q.p.n) != q.a.res {
			t.Errorf("输入：%v", q)
		}
	}
}
