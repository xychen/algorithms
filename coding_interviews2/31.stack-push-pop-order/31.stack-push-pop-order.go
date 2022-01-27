package problem31

// 剑指 Offer 31. 栈的压入、弹出序列

type Stack struct {
	stack []int
}

func (s *Stack) Push(x int) {
	s.stack = append(s.stack, x)
}
func (s *Stack) Pop() int {
	if len(s.stack) <= 0 {
		return -1
	}
	x := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return x
}
func (s *Stack) Top() int {
	if len(s.stack) <= 0 {
		return -1
	}
	return s.stack[len(s.stack)-1]
}

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) != len(popped) {
		return false
	}
	if len(pushed) == 0 && len(popped) == 0 {
		return true
	}
	s := &Stack{
		stack: make([]int, 0),
	}
	s.Push(pushed[0])
	i, j := 1, 0
	for i < len(pushed) || j < len(popped) {
		if s.Top() == popped[j] {
			s.Pop()
			j++
			continue
		}
		if i < len(pushed) {
			s.Push(pushed[i])
			i++
		} else {
			return false
		}
	}
	return true
}
