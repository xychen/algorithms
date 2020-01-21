package kit

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) PrintNodes() {
	fmt.Println(l.Vals())
}

func (l *ListNode) Vals() []int {
	p := l
	res := make([]int, 0)
	for p != nil {
		res = append(res, p.Val)
		p = p.Next
	}
	return res
}

//根据整型数组创建链表
func MakeListNode(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	res := &ListNode{}
	p := res
	for _, v := range vals {
		p.Next = &ListNode{Val: v}
		p = p.Next
	}
	return res.Next
}

//判断值是否相等
func IsValsEqual(l1 *ListNode, l2 *ListNode) bool {
	p1, p2 := l1, l2
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	//长度不同
	if p1 != nil || p2 != nil {
		return false
	}
	return true
}
