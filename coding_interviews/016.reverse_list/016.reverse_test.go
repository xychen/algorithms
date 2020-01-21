package problem016

import "testing"

import "algorithms/kit"

type param struct {
	l *ListNode
}

type ans struct {
	res *ListNode
}

type question struct {
	p param
	a ans
}

func Test_ReverseList(t *testing.T) {
	qs := []question{
		question{p: param{l: kit.MakeListNode([]int{1, 2, 3, 4})}, a: ans{res: kit.MakeListNode([]int{4, 3, 2, 1})}},
		question{p: param{l: kit.MakeListNode([]int{1})}, a: ans{res: kit.MakeListNode([]int{1})}},
		question{p: param{l: kit.MakeListNode([]int{})}, a: ans{res: kit.MakeListNode([]int{})}},
	}

	for i, q := range qs {
		res := ReverseList(q.p.l)
		if !kit.IsValsEqual(q.a.res, res) {
			t.Errorf("第 %d 组输入, %v", i, res)
		}
	}
}
