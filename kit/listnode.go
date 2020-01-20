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
