package problem017

import "algorithms/kit"

type ListNode = kit.ListNode

func MergeSortedList(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	p1, p2 := l1, l2
	res := &ListNode{}
	q := res
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			q.Next = p1
			p1 = p1.Next
		} else {
			q.Next = p2
			p2 = p2.Next
		}

		q = q.Next
	}
	if p1 != nil {
		q.Next = p1
	}
	if p2 != nil {
		q.Next = p2
	}
	return res.Next
}
