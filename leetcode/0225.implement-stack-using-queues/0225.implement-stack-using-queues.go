package problem0225

// 225. 用队列实现栈

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{
		queue: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop() int {
	n := len(this.queue)
	if n == 0 {
		return -1
	}
	x := this.queue[0]
	for n >= 1 {
		x = this.queue[0]
		this.queue = this.queue[1:]
		n--
		if n == 0 {
			continue
		}
		this.queue = append(this.queue, x)
	}
	return x
}

func (this *MyStack) Top() int {
	x := this.Pop()
	if x == -1 {
		return x
	}
	this.Push(x)
	return x
}

func (this *MyStack) Empty() bool {
	if len(this.queue) == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
