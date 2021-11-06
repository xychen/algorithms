package problem0203

// 203. 移除链表元素
// 给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode{}
	p := newHead
	for head != nil {
		if head.Val == val {
			head = head.Next
			continue
		}
		p.Next = head
		head = head.Next
		p = p.Next
	}
	p.Next = nil
	return newHead.Next
}
