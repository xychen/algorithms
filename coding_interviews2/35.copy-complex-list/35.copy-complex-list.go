package problem35

// 剑指 Offer 35. 复杂链表的复制

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	// 复制一个链表
	p := head
	for p != nil {
		tmp := &Node{
			Val:  p.Val,
			Next: p.Next,
		}
		p.Next = tmp
		p = tmp.Next
	}
	// 设置random指针
	p = head
	for p != nil {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}
	// 拆分链表
	p = head
	newHead := &Node{}
	q := newHead
	for p != nil {
		tmp := p.Next
		p.Next = tmp.Next
		p = p.Next
		tmp.Next = nil
		q.Next = tmp
		q = tmp
	}
	return newHead.Next
}
