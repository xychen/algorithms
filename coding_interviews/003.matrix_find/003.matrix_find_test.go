package problem003

import "testing"

type param struct {
	matrix  []int
	rows    int
	columns int
	target  int
}

type ans struct {
	res bool
}

type question struct {
	p param
	a ans
}

func Test_OK(t *testing.T) {
	m1 := []int{
		1, 2, 8, 9,
		2, 4, 9, 12,
		4, 7, 10, 13,
		6, 8, 11, 15,
	}

	qs := []question{
		question{p: param{m1, 4, 4, 16}, a: ans{res: false}},
		question{p: param{m1, 4, 4, 1}, a: ans{res: true}},
		question{p: param{m1, 4, 4, 6}, a: ans{res: true}},
		question{p: param{m1, 4, 4, 15}, a: ans{res: true}},
		question{p: param{m1, 4, 4, 9}, a: ans{res: true}},
	}

	for _, q := range qs {
		if q.a.res != Find(q.p.matrix, q.p.rows, q.p.columns, q.p.target) {
			t.Errorf("输入%v", q)
		}
	}
}
