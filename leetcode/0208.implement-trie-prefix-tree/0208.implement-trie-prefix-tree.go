package problem0208

type Trie struct {
	isEnd    bool
	children map[byte]*Trie
}

func Constructor() Trie {
	return Trie{
		isEnd:    false,
		children: make(map[byte]*Trie),
	}
}

func (this *Trie) Insert(word string) {
	node := this
	for i := 0; i < len(word); i++ {
		c := word[i]
		if _, ok := node.children[c]; !ok {
			nextNode := &Trie{
				isEnd:    false,
				children: make(map[byte]*Trie),
			}
			node.children[c] = nextNode
		}
		node = node.children[c]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for i := 0; i < len(word); i++ {
		c := word[i]
		if _, ok := node.children[c]; !ok {
			return false
		}
		node = node.children[c]
	}
	if node.isEnd {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for i := 0; i < len(prefix); i++ {
		c := prefix[i]
		if _, ok := node.children[c]; !ok {
			return false
		}
		node = node.children[c]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
