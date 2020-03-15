package problem021

import "algorithms/kit"

type StackWithMin struct {
	Stack    *kit.Stack
	MinStack *kit.Stack
}

func NewStackWithMin() *StackWithMin {
	return &StackWithMin{
		Stack:    kit.NewStack(),
		MinStack: kit.NewStack(),
	}
}

func (s *StackWithMin) Push(n int) {
	if !s.Stack.Empty() && s.Stack.Top() < n {
		s.MinStack.Push(s.Stack.Top())
	} else {
		s.MinStack.Push(n)
	}
	s.Stack.Push(n)
}

func (s *StackWithMin) Pop() int {
	s.MinStack.Pop()
	return s.Stack.Pop()
}

func (s *StackWithMin) Min() int {
	return s.MinStack.Top()
}
