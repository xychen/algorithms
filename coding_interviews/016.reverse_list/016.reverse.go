package problem016

import "algorithms/kit"

type ListNode = kit.ListNode

func ReverseList(l *ListNode) *ListNode {
	if l == nil {
		return nil
	}
	newh := &ListNode{}
	for l != nil {
		tmp := l
		l = l.Next
		tmp.Next = newh.Next
		newh.Next = tmp
	}
	return newh.Next
}
