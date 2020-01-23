package problem017

import "testing"

import "algorithms/kit"

type param struct {
	l1 *ListNode
	l2 *ListNode
}

type ans struct {
	res *ListNode
}

type question struct {
	p param
	a ans
}

func Test_MergeSortedList(t *testing.T) {
	qs := []question{
		question{p: param{l1: nil, l2: nil}, a: ans{res: nil}},
		question{p: param{l1: kit.MakeListNode([]int{1}), l2: nil}, a: ans{res: kit.MakeListNode([]int{1})}},
		question{p: param{l1: nil, l2: kit.MakeListNode([]int{1})}, a: ans{res: kit.MakeListNode([]int{1})}},
		question{p: param{l1: kit.MakeListNode([]int{1, 2}), l2: kit.MakeListNode([]int{3, 4})}, a: ans{res: kit.MakeListNode([]int{1, 2, 3, 4})}},
		question{p: param{l1: kit.MakeListNode([]int{3, 4}), l2: kit.MakeListNode([]int{1, 2})}, a: ans{res: kit.MakeListNode([]int{1, 2, 3, 4})}},
		question{p: param{l1: kit.MakeListNode([]int{1, 2, 3}), l2: kit.MakeListNode([]int{2, 4})}, a: ans{res: kit.MakeListNode([]int{1, 2, 2, 3, 4})}},
		question{p: param{l1: kit.MakeListNode([]int{4, 7}), l2: kit.MakeListNode([]int{1, 2, 3})}, a: ans{res: kit.MakeListNode([]int{1, 2, 3, 4, 7})}},
		question{p: param{l1: kit.MakeListNode([]int{1, 2, 5, 6}), l2: kit.MakeListNode([]int{3, 4})}, a: ans{res: kit.MakeListNode([]int{1, 2, 3, 4, 5, 6})}},
	}

	for i, q := range qs {
		res := MergeSortedList(q.p.l1, q.p.l2)
		if res == nil && q.a.res == nil {
			continue
		}
		if !kit.IsValsEqual(res, q.a.res) {
			t.Errorf("No.%d case cause error", i)
		}
	}
}
