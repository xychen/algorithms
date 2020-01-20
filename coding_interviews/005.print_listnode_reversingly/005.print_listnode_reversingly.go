package problem005

import (
	"algorithms/kit"
	"fmt"
)

type ListNode = kit.ListNode

//逆序打印链表：使用栈的方式
func PrintListReversingly_Iteratively(l *ListNode) {
	p := l
	stack := kit.NewStack()
	for p != nil {
		stack.Push(p.Val)
		p = p.Next
	}

	for !stack.Empty() {
		fmt.Printf("%d\t", stack.Pop())
	}
}

//逆序打印链表：递归方式
func PrintListReversingly_Recursively(l *ListNode) {
	if l != nil {
		if l.Next != nil {
			PrintListReversingly_Recursively(l.Next)
		}
		fmt.Printf("%d\t", l.Val)
	}

}
