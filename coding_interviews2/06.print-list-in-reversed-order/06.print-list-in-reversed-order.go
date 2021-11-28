package problem06

// 剑指 Offer 06. 从尾到头打印链表
// 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	p := head
	res := make([]int, 0)
	for p != nil {
		res = append(res, p.Val)
		p = p.Next
	}
	left, right := 0, len(res)-1
	for left <= right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
	return res
}
