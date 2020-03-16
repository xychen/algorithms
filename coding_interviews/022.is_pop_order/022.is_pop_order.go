package problem022

import "algorithms/kit"

//栈的压入、弹出顺序
func IsPopOrder(push []int, pop []int) bool {
	if len(push) != len(pop) {
		return false
	}
	if len(push) == 0 {
		return true
	}
	stack := kit.NewStack()
	for _, k := range push {
		stack.Push(k)
		for !stack.Empty() && stack.Top() == pop[0] {
			stack.Pop()
			pop = pop[1:]
		}
	}
	if stack.Empty() {
		return true
	}
	return false
}
