package problem0023

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodeHeap []*ListNode

func (l *ListNodeHeap) Len() int {
	return len(*l)
}

func (l *ListNodeHeap) Less(i, j int) bool {
	return (*l)[i].Val < (*l)[j].Val
}

func (l *ListNodeHeap) Swap(i, j int) {
	(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
}

// add x as element Len().
func (l *ListNodeHeap) Push(x interface{}) {
	*l = append(*l, x.(*ListNode))
}

// remove and return element Len() - 1.
func (l *ListNodeHeap) Pop() interface{} {
	//因为heap中的Pop函数会进行堆顶和最后一个元素的交换，并且调整堆，所以我们取最后一个元素
	old := *l
	n := len(old)
	x := old[n-1]
	*l = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}
	l := &ListNodeHeap{}
	heap.Init(l)
	//初始化
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(l, lists[i])
		}
	}
	res := &ListNode{}
	p := res
	for l.Len() > 0 {
		minNode := heap.Pop(l).(*ListNode)
		p.Next = minNode
		p = p.Next
		if minNode.Next != nil {
			heap.Push(l, minNode.Next)
		}
	}
	return res.Next
}
