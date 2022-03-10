package problem59_1

// 剑指 Offer 59 - I. 滑动窗口的最大值

type Dqueue struct {
	data []int
}

func (dq *Dqueue) PushBack(v int) {
	dq.data = append(dq.data, v)
}
func (dq *Dqueue) PopFront() int {
	v := dq.data[0]
	dq.data = dq.data[1:]
	return v
}
func (dq *Dqueue) PopBack() int {
	v := dq.data[len(dq.data)-1]
	dq.data = dq.data[:len(dq.data)-1]
	return v
}

func (dq *Dqueue) Back() int {
	return dq.data[len(dq.data)-1]
}

func (dq *Dqueue) Front() int {
	return dq.data[0]
}

func (dq *Dqueue) Empty() bool {
	return len(dq.data) == 0
}

func maxSlidingWindow(nums []int, k int) []int {
	if k < 1 || k > len(nums) {
		return []int{}
	}
	index := &Dqueue{
		data: make([]int, 0),
	}
	res := make([]int, 0)
	// 初始化
	for i := 0; i < k; i++ {
		for !index.Empty() && nums[i] > nums[index.Back()] {
			index.PopBack()
		}
		index.PushBack(i)
	}
	for i := k; i < len(nums); i++ {
		res = append(res, nums[index.Front()])
		for !index.Empty() && nums[i] > nums[index.Back()] {
			index.PopBack()
		}
		if !index.Empty() && (i-index.Front() >= k) {
			index.PopFront()
		}
		index.PushBack(i)
	}
	res = append(res, nums[index.Front()])
	return res
}
