package problem015

import "algorithms/kit"

type ListNode = kit.ListNode

func FindKthToTail(h *ListNode, k uint) *ListNode {
	if h == nil || k == 0 {
		return nil
	}
	fast, slow := h, h
	//有大于k个节点
	for i := 0; i < int(k)-1; i += 1 {
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
