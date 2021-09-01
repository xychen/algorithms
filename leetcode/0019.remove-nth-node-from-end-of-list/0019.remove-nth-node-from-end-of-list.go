package problem0019

//移除链表倒数第n个节点

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	lhead := &ListNode{}
	lhead.Next = head
	fast, slow := lhead, lhead

	for i := 0; i < n; i++ {
		fast = fast.Next
		if fast == nil {
			return head
		}
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return lhead.Next
}
