package problem18

// 剑指 Offer 18. 删除链表的节点

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{}
	newHead.Next = head
	p := newHead
	for p.Next != nil {
		if p.Next.Val == val {
			break
		}
		p = p.Next
	}
	if p.Next != nil {
		p.Next = p.Next.Next
	}
	return newHead.Next
}
