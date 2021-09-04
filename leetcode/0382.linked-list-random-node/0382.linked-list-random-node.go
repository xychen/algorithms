package problem0382

// 给定一个单链表，随机选择链表的一个节点，并返回相应的节点值。保证每个节点被选的概率一样。
// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247484974&idx=1&sn=795a33c338d4a5bd8d265bc7f9f63c03&chksm=9bd7f826aca07130e303d3d6f5c901b8aa00f9c3d02ffc26d45b56f1d36b538990c9eebd06a8&scene=21#wechat_redirect
// 当你遇到第i个元素时，应该有1/i的概率选择该元素，1 - 1/i的概率保持原有的选择 ， 1/i * (1 - 1/(i+1)) * （1 - /(i+2)）  =>  1/n

import "math/rand"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	Head *ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	return Solution{Head: head}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	if this.Head == nil {
		return -1
	}
	i := 0
	res := -1
	p := this.Head
	for p != nil {
		i++
		//生成[0, i)的随机数，随机数 = 0 的概率为， 1/i ，（i=1时肯定被选中，保证一定有返回值）
		if rand.Intn(i) == 0 {
			res = p.Val
		}
		p = p.Next
	}
	return res
}
