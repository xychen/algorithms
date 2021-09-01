package problem0160

//获取2个链表的交点
//方法一：同时走到尾结点，看尾结点是否一样。有相交点的话，长的先走差值步数，再一起走，找到交点

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	len1, len2 := 0, 0
	for p1.Next != nil {
		p1 = p1.Next
		len1++
	}
	for p2.Next != nil {
		p2 = p2.Next
		len2++
	}
	//最后一个节点不同，说明没有交点
	if p1 != p2 {
		return nil
	}
	p1, p2 = headA, headB
	if len1 > len2 {
		for i := 0; i < len1-len2; i++ {
			p1 = p1.Next
		}
	} else {
		for i := 0; i < len2-len1; i++ {
			p2 = p2.Next
		}
	}
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1

}
