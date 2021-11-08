package problem0138

// 138. 复制带随机指针的链表
// 给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random ，该指针可以指向链表中的任何节点或空节点。

// 构造这个链表的 深拷贝。

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
