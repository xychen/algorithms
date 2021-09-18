package problem0146

//lru

type DoubleLinkNode struct {
	key  int
	val  int
	pre  *DoubleLinkNode
	next *DoubleLinkNode
}

type LRUCache struct {
	capacity int
	size     int
	ht       map[int]*DoubleLinkNode
	head     *DoubleLinkNode
	tail     *DoubleLinkNode
}

func initDLinkNode(key, val int) *DoubleLinkNode {
	return &DoubleLinkNode{
		key: key,
		val: val,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		size:     0,
		ht:       make(map[int]*DoubleLinkNode),
		head:     initDLinkNode(0, 0),
		tail:     initDLinkNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.ht[key]; ok {
		this.MoveToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.ht[key]; ok {
		node.val = value
		this.MoveToHead(node)
		return
	}
	if this.size >= this.capacity {
		//删掉结尾的
		lastnode := this.tail.pre
		delete(this.ht, lastnode.key)
		this.DelNode(lastnode)
		this.size--
	}
	//添加新节点
	newnode := initDLinkNode(key, value)
	this.AddNodeToHead(newnode)
	this.ht[key] = newnode
	this.size++
}

func (this *LRUCache) MoveToHead(node *DoubleLinkNode) {
	this.DelNode(node)
	this.AddNodeToHead(node)
}

func (this *LRUCache) DelNode(node *DoubleLinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) AddNodeToHead(node *DoubleLinkNode) {
	node.next = this.head.next
	node.pre = this.head
	this.head.next = node
	node.next.pre = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
