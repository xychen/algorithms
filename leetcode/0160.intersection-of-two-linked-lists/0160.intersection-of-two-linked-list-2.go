package problem0160

//获取2个链表的交点
//方法二：p1走完链表1走链表2，p2走完链表2走链表1，找到交点，如果不想交，p1/p2都会走到null

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}

		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}
