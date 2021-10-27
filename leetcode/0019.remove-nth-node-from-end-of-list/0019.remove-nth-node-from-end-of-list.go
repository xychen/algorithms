package problem0019

//移除链表倒数第n个节点

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{}
	newHead.Next = head
	slow, fast := newHead, newHead
	// fast走n+1步
	for i := 0; i <= n; i++ {
		if fast == nil {
			return head
		}
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return newHead.Next
}
