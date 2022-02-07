package problem0239

// 239. 滑动窗口最大值
// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

// 返回 滑动窗口中的最大值 。

type myqueue struct {
	data []int
}

func NewQue() *myqueue {
	return &myqueue{
		data: make([]int, 0),
	}
}

func (que *myqueue) push(x int) {
	for len(que.data) > 0 && que.back() < x {
		que.data = que.data[:len(que.data)-1]
	}
	que.data = append(que.data, x)
}
func (que *myqueue) pop(x int) {
	if len(que.data) > 0 && que.front() == x {
		que.data = que.data[1:]
	}
}
func (que *myqueue) front() int {
	return que.data[0]
}

func (que *myqueue) back() int {
	return que.data[len(que.data)-1]
}

func maxSlidingWindow(nums []int, k int) []int {
	que := NewQue()
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			que.push(nums[i])
		} else {
			que.push(nums[i])
			res = append(res, que.front())
			que.pop(nums[i-k+1])
		}
	}
	return res
}
