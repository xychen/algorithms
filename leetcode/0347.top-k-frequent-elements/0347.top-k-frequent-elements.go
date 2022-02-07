package problem0347

// 347. 前 K 个高频元素

/**************使用快排partition思想********************/

type node struct {
	key int
	val int
}

func topKFrequent(nums []int, k int) []int {
	ht := make(map[int]int)
	for _, num := range nums {
		ht[num]++
	}
	nodeList := make([]node, len(ht))
	i := 0
	for k, v := range ht {
		nodeList[i] = node{key: k, val: v}
		i++
	}
	left, right := 0, len(nodeList)-1
	for {
		index := partition(nodeList, left, right)
		if index == k-1 {
			break
		}
		if index > k-1 {
			right = index - 1
		} else if index < k-1 {
			left = index + 1
		}
	}
	// fmt.Println(nodeList)
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = nodeList[i].key
	}
	return res
}

func partition(nodeList []node, left, right int) int {
	guardNode := nodeList[left]
	// 大的等于的放左边，小的放右边
	for left < right {
		for left < right && nodeList[right].val <= guardNode.val {
			right--
		}
		nodeList[left] = nodeList[right]
		for left < right && nodeList[left].val > guardNode.val {
			left++
		}
		nodeList[right] = nodeList[left]
	}
	nodeList[left] = guardNode
	return left
}

/**************小顶堆实现方式********************/
type heapNode struct {
	num int
	c   int
}

func topKFrequent2(nums []int, k int) []int {
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
