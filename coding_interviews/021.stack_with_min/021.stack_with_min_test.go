package problem021

import (
	"strconv"
	"strings"
	"testing"
)

type parma struct {
	elements []string
}

type ans struct {
	res int
}

type question struct {
	p parma
	a ans
}

func Test_StackWithMin(t *testing.T) {
	qs := []question{
		question{p: parma{elements: []string{"push_2", "push_1", "push_3"}}, a: ans{res: 1}},
		question{p: parma{elements: []string{"push_2", "push_1", "push_3", "pop_"}}, a: ans{res: 1}},
		question{p: parma{elements: []string{"push_2", "push_1", "push_3", "pop_", "pop_"}}, a: ans{res: 2}},
	}

	for i, q := range qs {
		s := NewStackWithMin()
		for _, element := range q.p.elements {
			info := strings.Split(element, "_")
			if "push" == info[0] {
				n, _ := strconv.Atoi(info[1])
				s.Push(n)
			} else if "pop" == info[0] {
				s.Pop()
			}
		}
		if s.Min() != q.a.res {
			t.Errorf("question %d is error", i)
		}
	}
}
