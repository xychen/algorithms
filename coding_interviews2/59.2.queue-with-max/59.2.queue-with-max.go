package problem59_2

// 剑指 Offer 59 - II. 队列的最大值

type Element struct {
	val int
	idx int
}

type MaxQueue struct {
	dataQ   []Element
	maxValQ []Element
	idx     int
}

func Constructor() MaxQueue {
	return MaxQueue{
		dataQ:   make([]Element, 0),
		maxValQ: make([]Element, 0),
		idx:     0,
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.maxValQ) > 0 {
		return this.maxValQ[0].val
	}
	return -1
}

func (this *MaxQueue) Push_back(value int) {
	for len(this.maxValQ) > 0 && value > this.maxValQ[len(this.maxValQ)-1].val {
		this.maxValQ = this.maxValQ[:len(this.maxValQ)-1] //pop_back
	}
	element := Element{
		val: value,
		idx: this.idx,
	}
	this.idx++
	this.dataQ = append(this.dataQ, element)
	this.maxValQ = append(this.maxValQ, element)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.dataQ) <= 0 {
		return -1
	}
	element := this.dataQ[0]
	this.dataQ = this.dataQ[1:]
	if element.idx == this.maxValQ[0].idx {
		this.maxValQ = this.maxValQ[1:]
	}
	return element.val
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
