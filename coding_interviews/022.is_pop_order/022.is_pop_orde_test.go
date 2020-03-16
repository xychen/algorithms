package problem022

import "testing"

type parma struct {
	push []int
	pop  []int
}

type ans struct {
	res bool
}

type question struct {
	p parma
	a ans
}

func Test_IsPopOrder(t *testing.T) {
	qs := []question{
		question{p: parma{push: []int{1, 2, 3, 4, 5}, pop: []int{4, 3, 5, 1, 2}}, a: ans{res: false}},
		question{p: parma{push: []int{1, 2, 3, 4, 5}, pop: []int{5, 4, 3, 2, 1}}, a: ans{res: true}},
		question{p: parma{push: []int{1, 2, 3, 4, 5}, pop: []int{4, 5, 3, 2, 1}}, a: ans{res: true}},
		question{p: parma{push: []int{1, 2, 3, 4, 5}, pop: []int{1, 2, 3, 4, 5}}, a: ans{res: true}},
		question{p: parma{push: []int{}, pop: []int{}}, a: ans{res: true}},
		question{p: parma{push: []int{1, 2}, pop: []int{1}}, a: ans{res: false}},
	}
	for i, q := range qs {
		if IsPopOrder(q.p.push, q.p.pop) != q.a.res {
			t.Errorf("case No.%d is error", i)
		}
	}
}
