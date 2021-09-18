package problem0496

// 下一个更大元素 I

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

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	s := initStack()
	ansMap := make(map[int]int)
	for i := len(nums2) - 1; i >= 0; i-- {
		for !s.Empty() && s.Top() <= nums2[i] {
			s.Pop()
		}
		if s.Empty() {
			ansMap[nums2[i]] = -1
		} else {
			ansMap[nums2[i]] = s.Top()
		}
		s.Push(nums2[i])
	}

	ans := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		ans[i] = ansMap[nums1[i]]
	}
	return ans
}
