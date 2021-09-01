package problem0141

//判断链表是否有环

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head
	for fast != nil {
		//fast走2步
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		//slow走一步
		slow = slow.Next
		if fast != nil && fast == slow {
			return true
		}
	}
	return false
}
