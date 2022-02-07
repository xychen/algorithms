package problem0503

// 下一个更大元素 II

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

func nextGreaterElements(nums []int) []int {
	ans := make([]int, len(nums))
	n := len(nums)
	s := initStack()
	for i := 2*n - 1; i >= 0; i-- {
		for !s.Empty() && s.Top() <= nums[i%n] {
			s.Pop()
		}
		if s.Empty() {
			ans[i%n] = -1
		} else {
			ans[i%n] = s.Top()
		}
		s.Push(nums[i%n])
	}
	return ans
}

func nextGreaterElements2(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	stack := make([]int, 0)
	for i := 2*n - 1; i >= 0; i-- {
		// realIndex := i % n
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i%n] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ans[i%n] = stack[len(stack)-1]
		} else {
			ans[i%n] = -1
		}
		stack = append(stack, nums[i%n])
	}
	return ans
}
