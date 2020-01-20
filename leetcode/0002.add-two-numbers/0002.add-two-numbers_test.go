package problem0002

import (
	"testing"
)

type params struct {
	one *ListNode
	two *ListNode
}

type ans struct {
	res *ListNode
}

type question struct {
	p params
	a ans
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

func listEqual(one *ListNode, two *ListNode) bool {
	for one != nil && two != nil {
		if one.Val != two.Val {
			return false
		}
		one = one.Next
		two = two.Next
	}

	if (one == nil && two != nil) || (one != nil && two == nil) {
		return false
	}

	return true
}

func Test_OK(t *testing.T) {
	qs := []question{
		question{
			p: params{
				one: makeListNode([]int{2, 4, 3}),
				two: makeListNode([]int{5, 6, 4}),
			},
			a: ans{
				res: makeListNode([]int{7, 0, 8}),
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		if !listEqual(a.res, addTwoNumbers(p.one, p.two)) {
			t.Errorf("输入: %v, %v", p.one.Vals(), p.two.Vals())
		}
	}
}
