package problem0024

// 24. 两两交换链表中的节点

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{}
	p := newHead
	slow, fast := head, head.Next
	for slow != nil {
		if fast == nil {
			p.Next = slow
			break
		}
		t1 := slow
		t1.Next = nil
		t2 := fast
		slow = fast.Next
		if slow != nil {
			fast = slow.Next
		} else {
			fast = nil
		}
		t2.Next = t1
		p.Next = t2
		p = t1
	}
	return newHead.Next
}
