package problem015

import "testing"

type param struct {
	l *ListNode
	k uint
}
type ans struct {
	res int
}

type question struct {
	p param
	a *ans
}

func Test_FindKthToTail(t *testing.T) {
	qs := []question{
		question{p: param{l: makeListNode([]int{1, 2, 3, 4, 5}), k: 2}, a: &ans{4}},
		question{p: param{l: makeListNode([]int{1, 2}), k: 2}, a: &ans{1}},
		question{p: param{l: makeListNode([]int{1, 2}), k: 1}, a: &ans{2}},
		question{p: param{l: makeListNode([]int{1, 2}), k: 3}, a: nil},
	}

	for i, q := range qs {
		res := FindKthToTail(q.p.l, q.p.k)
		if res == nil {
			if q.a != nil {
				t.Errorf("第 %d 组输入, %v", i, res)
			}
		} else {
			if (q.a == nil) || q.a.res != res.Val {
				t.Errorf("第 %d 组输入, %v", i, res)
			}
		}
	}
}

func makeListNode(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	res := &ListNode{}
	p := res
	for _, v := range vals {
		p.Next = &ListNode{Val: v}
		p = p.Next
	}
	return res.Next
}
