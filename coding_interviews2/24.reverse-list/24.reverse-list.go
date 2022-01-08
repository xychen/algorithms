package problem24

// 剑指 Offer 24. 反转链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{}
	p := head
	for p != nil {
		tmp := p
		p = p.Next
		tmp.Next = newHead.Next
		newHead.Next = tmp
	}
	return newHead.Next
}
