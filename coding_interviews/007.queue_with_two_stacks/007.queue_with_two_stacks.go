package problem007

import (
	"algorithms/kit"
	"errors"
)

type Stack = kit.Stack

type CQueue struct {
	s1 *Stack
	s2 *Stack
}

func NewCQueue() *CQueue {
	return &CQueue{s1: kit.NewStack(), s2: kit.NewStack()}
}

//追加到尾部
func (cq *CQueue) AppendTail(val int) {
	cq.s1.Push(val)
}

//队列头部删除
func (cq *CQueue) DeleteHead() (int, error) {
	//如果s2为空
	if cq.s2.Empty() {
		for !cq.s1.Empty() {
			n := cq.s1.Pop()
			cq.s2.Push(n)
		}
	}
	if cq.s2.Empty() {
		return -1, errors.New("empty queue")
	}
	return cq.s2.Pop(), nil
}
