package problem020

import (
	"testing"
)

type parma struct {
	numbers []int
	columns int
	rows    int
}

type ans struct {
	res []int
}

type question struct {
	p parma
	a ans
}

func Test_PrintMatrixInCircle(t *testing.T) {
	qs := []question{
		question{p: parma{numbers: []int{1}, columns: 1, rows: 1}, a: ans{res: []int{1}}},
		question{p: parma{numbers: []int{1, 2}, columns: 2, rows: 1}, a: ans{res: []int{1, 2}}},
		question{p: parma{numbers: []int{1, 2, 3, 4, 5, 6}, columns: 3, rows: 2}, a: ans{res: []int{1, 2, 3, 6, 5, 4}}},
		question{p: parma{numbers: []int{1, 2, 3, 4, 5, 6}, columns: 2, rows: 3}, a: ans{res: []int{1, 2, 4, 6, 5, 3}}},
		question{p: parma{numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, columns: 3, rows: 3}, a: ans{res: []int{1, 2, 3, 6, 9, 8, 7, 4}}},
	}

	for qi, q := range qs {
		res := PrintMatrixInCircle(q.p.numbers, q.p.columns, q.p.rows, 0)
		if len(q.a.res) != len(res) {
			t.Errorf("PrintMatrixInCircle: %d len error", qi)
			continue
		}
		for i, k := range res {
			if k != q.a.res[i] {
				t.Errorf("PrintMatrixInCircle: %d error", qi)
			}
		}
	}
}

func Test_PrintMatrixClockwisely(t *testing.T) {
	qs := []question{
		question{p: parma{numbers: []int{1}, columns: 1, rows: 1}, a: ans{res: []int{1}}},
		question{p: parma{numbers: []int{1, 2}, columns: 2, rows: 1}, a: ans{res: []int{1, 2}}},
		question{p: parma{numbers: []int{1, 2, 3, 4, 5, 6}, columns: 3, rows: 2}, a: ans{res: []int{1, 2, 3, 6, 5, 4}}},
		question{p: parma{numbers: []int{1, 2, 3, 4, 5, 6}, columns: 2, rows: 3}, a: ans{res: []int{1, 2, 4, 6, 5, 3}}},
		question{p: parma{numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, columns: 3, rows: 3}, a: ans{res: []int{1, 2, 3, 6, 9, 8, 7, 4, 5}}},
		question{p: parma{numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, columns: 4, rows: 4}, a: ans{res: []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10}}},
		question{p: parma{numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}, columns: 5, rows: 5}, a: ans{res: []int{1, 2, 3, 4, 5, 10, 15, 20, 25, 24, 23, 22, 21, 16, 11, 6, 7, 8, 9, 14, 19, 18, 17, 12, 13}}},
	}

	for qi, q := range qs {
		res := PrintMatrixClockwisely(q.p.numbers, q.p.columns, q.p.rows)
		if len(q.a.res) != len(res) {
			t.Errorf("PrintMatrixClockwisely: %d len error", qi)
			continue
		}
		for i, k := range res {
			if k != q.a.res[i] {
				t.Errorf("PrintMatrixClockwisely: %d error", qi)
			}
		}
	}
}
