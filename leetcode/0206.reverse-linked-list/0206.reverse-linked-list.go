package problem0206

// 206. 反转链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	newHead := &ListNode{}
	for head != nil {
		cur := head
		head = head.Next
		cur.Next = newHead.Next
		newHead.Next = cur
	}
	return newHead.Next
}
