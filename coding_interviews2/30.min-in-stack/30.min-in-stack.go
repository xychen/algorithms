package problem30

// 剑指 Offer 30. 包含min函数的栈

type MinStack struct {
	stack    []int
	minstack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minstack: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	if len(this.stack) <= 0 {
		this.minstack = append(this.minstack, x)
	} else {
		t := this.minstack[len(this.minstack)-1]
		if x < t {
			t = x
		}
		this.minstack = append(this.minstack, t)
	}
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	if len(this.stack) <= 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
	this.minstack = this.minstack[:len(this.minstack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) <= 0 {
		return -1
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	if len(this.minstack) <= 0 {
		return -1
	}
	return this.minstack[len(this.minstack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
