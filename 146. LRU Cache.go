package leetcode

type LRUCache struct {
	cap      int
	data     map[int]*Node
	beg, end *Node
}

type Node struct {
	key       int
	value     int
	pre, next *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, make(map[int]*Node), nil, nil}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.data[key]
	if !ok {
		return -1
	}
	this.putNodeAhead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.data[key]
	if ok {
		node.value = value
		this.putNodeAhead(node)
	} else {
		if len(this.data) >= this.cap {
			// 容量满，删除末尾节点
			this.delNodeEnd()
		}
		// 新节点插在头部
		node := &Node{key, value, nil, this.beg}
		if this.beg != nil {
			this.beg.pre = node
		}
		this.beg = node
		if node.next == nil {
			this.end = node
		}
		this.data[key] = node
	}
}

// 将节点提前
func (this *LRUCache) putNodeAhead(node *Node) {
	if node.pre == nil {
		return
	}

	node.pre.next = node.next
	if node.next != nil {
		node.next.pre = node.pre
	} else {
		this.end = node.pre
	}

	node.next = this.beg
	node.pre = nil
	if this.beg != nil {
		this.beg.pre = node
	}
	this.beg = node
}

// 删除末尾节点
func (this *LRUCache) delNodeEnd() {
	node := this.end

	if node.pre == nil {
		this.beg = nil
	} else {
		node.pre.next = nil
	}
	this.end = node.pre
	delete(this.data, node.key)
}
