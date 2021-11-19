package problem0347

// 347. 前 K 个高频元素

type heapNode struct {
	num int
	c   int
}

func topKFrequent(nums []int, k int) []int {
	ht := make(map[int]int)
	for _, num := range nums {
		ht[num]++
	}
	heap := make([]heapNode, k)
	i := 0
	for num, c := range ht {
		node := heapNode{
			num: num,
			c:   c,
		}
		if i < k-1 {
			heap[i] = node
		} else if i == k-1 {
			// 前k个元素构建堆
			heap[i] = node
			buildHeap(heap)
		} else {
			if c > heap[0].c {
				// 替换堆顶元素
				heap[0] = node
				// 调整堆
				ajust(heap, 0)
			}
		}
		i++
	}
	res := make([]int, k)
	for i, node := range heap {
		res[i] = node.num
	}
	return res
}

func ajust(heap []heapNode, start int) {
	// 自上而下调整堆
	i, n := start, len(heap)
	for i < n {
		childIndex := 2*i + 1
		if 2*(i+1) < n && heap[2*(i+1)].c < heap[2*i+1].c {
			childIndex = 2 * (i + 1)
		}
		if childIndex >= n {
			break
		}
		// 较小的元素调上来
		if heap[childIndex].c < heap[i].c {
			heap[i], heap[childIndex] = heap[childIndex], heap[i]
		} else {
			break
		}
		i = childIndex
	}
}

// 自下而上构建小顶堆
func buildHeap(heap []heapNode) {
	n := len(heap)
	// fmt.Println(heap)
	for i := (n - 1) / 2; i >= 0; i-- {
		ajust(heap, i)
	}
	// fmt.Println(heap)
}
