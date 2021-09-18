package problem0460

// LFU

import "fmt"

type Node struct {
	key   int
	value int
	freq  int
	pre   *Node
	next  *Node
}

type DoubleLink struct {
	head *Node
	tail *Node
}

type LFUCache struct {
	capacity  int
	freqMapDL map[int]*DoubleLink
	kv        map[int]*Node
	minFreq   int
}

func initNode(key, value, freq int) *Node {
	return &Node{
		key:   key,
		value: value,
		freq:  freq,
	}
}

func initDoubleLink() *DoubleLink {
	dl := &DoubleLink{
		head: initNode(0, 0, 0),
		tail: initNode(0, 0, 0),
	}
	dl.head.next = dl.tail
	dl.tail.pre = dl.head
	return dl
}

func Constructor(capacity int) LFUCache {
	lfu := LFUCache{
		capacity:  capacity,
		freqMapDL: make(map[int]*DoubleLink),
		kv:        make(map[int]*Node),
		minFreq:   1,
	}
	return lfu
}

func (this *LFUCache) Get(key int) int {
	if node, ok := this.kv[key]; ok {
		this.upgradeFreq(key)
		return node.value
	}
	return -1
}

func (this *LFUCache) upgradeFreq(key int) {
	node := this.kv[key]
	//原来链表中删除该节点
	node.pre.next = node.next
	node.next.pre = node.pre
	dl := this.freqMapDL[node.freq]
	//如果链表空了
	if dl.head.next == dl.tail {
		delete(this.freqMapDL, node.freq)
		if this.minFreq == node.freq {
			this.minFreq++
		}
	}
	node.freq++
	this.addNodeToDLHead(node)
}

func (this *LFUCache) addNodeToDLHead(node *Node) {
	//新创建链表
	if _, ok := this.freqMapDL[node.freq]; !ok {
		this.freqMapDL[node.freq] = initDoubleLink()
	}
	dl := this.freqMapDL[node.freq]
	//加入到链表中(头插)
	node.next = dl.head.next
	node.pre = dl.head
	node.next.pre = node
	dl.head.next = node
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity <= 0 {
		return
	}
	if node, ok := this.kv[key]; ok {
		node.value = value
		this.upgradeFreq(key)
		return
	}

	if len(this.kv) >= this.capacity {
		this.removeMinFreqNode()
	}
	//添加新节点
	node := initNode(key, value, 1)
	this.addNodeToDLHead(node)
	this.kv[key] = node
	this.minFreq = 1
}

func (this *LFUCache) removeMinFreqNode() {
	dl := this.freqMapDL[this.minFreq]
	// show(dl)
	//从尾部删除
	lastnode := dl.tail.pre
	// fmt.Println(dl, this.minFreq, lastnode)
	lastnode.pre.next = lastnode.next
	lastnode.next.pre = lastnode.pre
	//该频次没有数据了
	if dl.head.next == dl.tail {
		delete(this.freqMapDL, this.minFreq)
		//不需要更新 minFreq，因为put的时候才可能引起removeMinFreqNode，此时minFreq会设置成1
	}
	delete(this.kv, lastnode.key)
}

//调试用
func show(dl *DoubleLink) {
	p := dl.head
	for p != nil {
		fmt.Printf("%d\t", p.value)
		p = p.next
	}
	fmt.Println()
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
