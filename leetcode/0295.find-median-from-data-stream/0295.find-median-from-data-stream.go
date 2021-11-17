package problem0295

import (
	"container/heap"
	"sort"
)

// 295. 数据流的中位数

// 使用官方的堆实现
type IntHeap struct {
	sort.IntSlice
}

func (hp *IntHeap) Push(v interface{}) {
	hp.IntSlice = append(hp.IntSlice, v.(int))
}

func (hp *IntHeap) Pop() interface{} {
	n := len(hp.IntSlice)
	v := hp.IntSlice[n-1]
	hp.IntSlice = hp.IntSlice[:n-1]
	return v
}

type MedianFinder struct {
	leftHeap  IntHeap
	rightHeap IntHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		leftHeap:  IntHeap{},
		rightHeap: IntHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	n := this.leftHeap.Len() + this.rightHeap.Len()
	if n&1 == 0 {
		// 插入到左边（大顶堆，负数的小顶堆）  => 先插到右边，pop出来再取反插入到左边
		heap.Push(&this.rightHeap, num)
		heap.Push(&this.leftHeap, -heap.Pop(&this.rightHeap).(int))
	} else {
		// 插入到右边（小顶堆）  => 先取反插到左边，pop出来再取反插入到右边
		heap.Push(&this.leftHeap, -num)
		heap.Push(&this.rightHeap, -heap.Pop(&this.leftHeap).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	n := this.leftHeap.Len() + this.rightHeap.Len()
	if n == 0 {
		return float64(0)
	}
	if n&1 == 0 {
		return (float64(-this.leftHeap.IntSlice[0]) + float64(this.rightHeap.IntSlice[0])) / 2
	} else {
		return float64(-this.leftHeap.IntSlice[0])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

// 方法二：自己实现堆操作

type MedianFinder2 struct {
	leftHeap  []int //大顶堆
	rightHeap []int //小顶堆
}

func Constructor2() MedianFinder2 {
	return MedianFinder2{
		leftHeap:  make([]int, 0),
		rightHeap: make([]int, 0),
	}
}

func (this *MedianFinder2) AddNum(num int) {
	n := len(this.leftHeap) + len(this.rightHeap)
	if n%2 == 0 {
		// 插入到左堆中 => 先插右堆，从右堆中找到最小的，插入到左堆
		this.rightHeap = append(this.rightHeap, num)
		this.HeapSwim(this.rightHeap, len(this.rightHeap)-1, true)
		v := this.rightHeap[0]
		this.rightHeap[0] = this.rightHeap[len(this.rightHeap)-1]
		this.rightHeap = this.rightHeap[:len(this.rightHeap)-1]
		this.HeapSink(this.rightHeap, 0, true)

		this.leftHeap = append(this.leftHeap, v)
		this.HeapSwim(this.leftHeap, len(this.leftHeap)-1, false)

	} else {
		// 插入到右堆中 => 先插左堆，从左堆中找到最大的，插入到右堆
		this.leftHeap = append(this.leftHeap, num)
		this.HeapSwim(this.leftHeap, len(this.leftHeap)-1, false)
		v := this.leftHeap[0]
		this.leftHeap[0] = this.leftHeap[len(this.leftHeap)-1]
		this.leftHeap = this.leftHeap[:len(this.leftHeap)-1]
		this.HeapSink(this.leftHeap, 0, false)

		this.rightHeap = append(this.rightHeap, v)
		this.HeapSwim(this.rightHeap, len(this.rightHeap)-1, true)

	}
}

func (this *MedianFinder2) HeapSink(h []int, k int, isMinHeap bool) {
	if k < 0 {
		return
	}
	n := len(h)
	for 2*k+1 < n {
		child := 2*k + 1
		// 小顶堆
		if isMinHeap {
			if child+1 < n && h[child+1] < h[child] {
				child += 1
			}
			if child < n && h[child] < h[k] {
				h[k], h[child] = h[child], h[k]
			}
		} else {
			// 大顶堆
			if child+1 < n && h[child+1] > h[child] {
				child += 1
			}
			if child < n && h[child] > h[k] {
				h[k], h[child] = h[child], h[k]
			}
		}
		k = child
	}
}

func (this *MedianFinder2) HeapSwim(h []int, k int, isMinHeap bool) {
	if k >= len(h) {
		return
	}
	for (k-1)/2 >= 0 {
		parent := (k - 1) / 2
		// 小顶堆，小元素上浮
		if isMinHeap {
			if h[k] < h[parent] {
				h[k], h[parent] = h[parent], h[k]
			} else {
				break
			}
		} else {
			if h[k] > h[parent] {
				h[k], h[parent] = h[parent], h[k]
			} else {
				break
			}
		}
		k = parent
	}
}

func (this *MedianFinder2) FindMedian() float64 {
	n := len(this.leftHeap) + len(this.rightHeap)
	if n == 0 {
		return float64(0)
	} else if n%2 == 0 {
		return float64(this.leftHeap[0]+this.rightHeap[0]) / 2
	}
	return float64(this.leftHeap[0])
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
