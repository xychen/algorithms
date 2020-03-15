package kit

type Stack struct {
	nums []int
}

func NewStack() *Stack {
	return &Stack{nums: []int{}}
}

func (s *Stack) Push(n int) {
	s.nums = append(s.nums, n)
}

func (s *Stack) Pop() int {
	n := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return n
}

func (s *Stack) Len() int {
	return len(s.nums)
}

func (s *Stack) Empty() bool {
	return len(s.nums) == 0
}

func (s *Stack) Top() int {
	return s.nums[len(s.nums)-1]
}
