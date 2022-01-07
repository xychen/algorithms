package problem22

// 剑指 Offer 22. 链表中倒数第k个节点

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	p1, p2 := head, head
	for k > 0 && p2 != nil {
		p2 = p2.Next
		k--
	}
	if k > 0 {
		return nil
	}

	for p2 != nil {
		p2 = p2.Next
		p1 = p1.Next
	}
	return p1
}
