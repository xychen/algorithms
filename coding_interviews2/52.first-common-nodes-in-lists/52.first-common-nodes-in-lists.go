package problem52

type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法一：巧妙的办法，见题解：https://leetcode-cn.com/problems/intersection-of-two-linked-lists/solution/xiang-jiao-lian-biao-by-leetcode-solutio-a8jn/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
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

// 方法二：先计算长度，再求交点
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headB == nil || headA == nil {
		return nil
	}
	p1, p2 := headA, headB
	c1, c2 := 1, 1
	for p1.Next != nil || p2.Next != nil {
		if p1.Next != nil {
			p1 = p1.Next
			c1++
		}
		if p2.Next != nil {
			p2 = p2.Next
			c2++
		}
	}
	if p1 != p2 {
		return nil
	}
	diff := c1 - c2
	p1, p2 = headA, headB
	if c2 > c1 {
		p1, p2 = headB, headA
		diff = -diff
	}
	for diff > 0 {
		p1 = p1.Next
		diff--
	}
	for {
		if p1 == p2 {
			return p1
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return nil
}
