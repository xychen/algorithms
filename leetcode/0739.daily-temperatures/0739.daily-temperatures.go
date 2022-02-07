package problem0739

// 每日温度

func dailyTemperatures(temperatures []int) []int {
	s := initStack()
	ans := make([]int, len(temperatures))
	for i := len(temperatures) - 1; i >= 0; i-- {
		for !s.Empty() && temperatures[s.Top()] <= temperatures[i] {
			s.Pop()
		}
		if s.Empty() {
			ans[i] = 0
		} else {
			ans[i] = s.Top() - i
		}
		s.Push(i)
	}
	return ans
}

type stack struct {
	stack []int
}

func initStack() *stack {
	return &stack{
		stack: make([]int, 0),
	}
}

func (s *stack) Empty() bool {
	if len(s.stack) == 0 {
		return true
	}
	return false
}

func (s *stack) Push(x int) {
	s.stack = append(s.stack, x)
}

func (s *stack) Pop() int {
	l := len(s.stack)
	x := s.stack[l-1]
	s.stack = s.stack[0 : l-1]
	return x
}

func (s *stack) Top() int {
	return s.stack[len(s.stack)-1]
}

func dailyTemperatures2(temperatures []int) []int {
	stack := make([]int, 0, len(temperatures))
	ans := make([]int, len(temperatures))
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return ans
}
