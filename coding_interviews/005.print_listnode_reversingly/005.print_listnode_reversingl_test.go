package problem005

import "testing"

import "fmt"

func makeListNode(vals []int) *ListNode {
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

func Test_PrintListReversingly_Iteratively(t *testing.T) {
	l := makeListNode([]int{1, 2, 3, 4, 5})
	PrintListReversingly_Iteratively(l)
	fmt.Println()
}

func Test_PrintListReversingly_Recursively(t *testing.T) {
	l := makeListNode([]int{1, 2, 3, 4, 5})
	PrintListReversingly_Recursively(l)
	fmt.Println()
}
