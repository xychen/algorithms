package problem09

// 剑指 Offer 09. 用两个栈实现队列
// 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

type CQueue struct {
	// 队尾
	stack1 []int
	// 队头
	stack2 []int
}

func Constructor() CQueue {
	return CQueue{
		stack1: make([]int, 0),
		stack2: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1 = append(this.stack1, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.stack1) == 0 && len(this.stack2) == 0 {
		return -1
	}
	if len(this.stack2) == 0 {
		for len(this.stack1) > 0 {
			x := this.stack1[len(this.stack1)-1]
			this.stack1 = this.stack1[0 : len(this.stack1)-1]
			this.stack2 = append(this.stack2, x)
		}
	}
	x := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[0 : len(this.stack2)-1]
	return x
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
