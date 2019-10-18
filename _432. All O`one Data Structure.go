package leetcode

type AllOne struct {
	keyNodeMap   map[string]*Node
	valueNodeMap map[int]*Node
	beg          *Node // maxNode
	end          *Node // minNode
}

type Node struct {
	value int
	keys  map[string]struct{} // 存放值为value的key
	pre   *Node
	next  *Node
}

func Constructor() AllOne {
	return AllOne{make(map[string]*Node), make(map[int]*Node), nil, nil}
}

func (this *AllOne) Inc(key string) {
	node, ok := this.keyNodeMap[key]
	if !ok {
		// key不存在，插在尾部value为1的node，如果不存在，则创建
		if this.end == nil {
			// 链表为空，创建value为1的node
			newNode := &Node{1, map[string]struct{}{key: {}}, nil, nil}
			this.keyNodeMap[key], this.beg, this.end = newNode, newNode, newNode
		} else if this.end.value != 1 {
			// 没有value为1的node，创建value为1的node
			newNode := &Node{1, map[string]struct{}{key: {}}, this.end, nil}
			this.keyNodeMap[key], this.beg, this.end, this.end.next = newNode, newNode, newNode, newNode
		} else {
			// 存在value为1的node
			this.end.keys[key] = struct{}{}
			this.keyNodeMap[key] = this.end
		}
		return
	}

	// key存在, 将该key添加进value+1的node
	if node.pre == nil {
		newNode := &Node{node.value + 1, map[string]struct{}{key: {}}, nil, node}
		this.beg, node.pre, this.keyNodeMap[key] = newNode, newNode, newNode
	} else if node.pre.value != node.value+1 {
		newNode := &Node{node.value + 1, map[string]struct{}{key: {}}, node.pre, node}
		node.pre, node.pre.next, this.keyNodeMap[key] = newNode, newNode, newNode
	} else {
		node.pre.keys[key] = struct{}{}
		this.keyNodeMap[key] = node.pre
	}

	// 从value的node中删除
	delete(node.keys, key)
	if len(node.keys) == 0 {
		// 没有key的值为node.value, 删除该节点
		if node.pre != nil {
			node.pre.next = node.next
		}
		if node.next != nil {
			node.next.pre = node.pre
		}

		if this.beg == node {
			this.beg = node.pre
		}
		if this.end == node {
			this.end = node.pre
		}
	}
}

func (this *AllOne) Dec(key string) {
	node, ok := this.keyNodeMap[key]
	if !ok {
		return
	}
	// key存在, 将该key添加进value-1的node
	if node.value != 1 {
		if node.next == nil {
			newNode := &Node{node.value - 1, map[string]struct{}{key: {}}, node, nil}
			this.end, node.next, this.keyNodeMap[key] = newNode, newNode, newNode
		} else if node.next.value != node.value-1 {
			newNode := &Node{node.value - 1, map[string]struct{}{key: {}}, node, node.next}
			node.next, node.next.pre, this.keyNodeMap[key] = newNode, newNode, newNode
		} else {
			node.next.keys[key] = struct{}{}
			this.keyNodeMap[key] = node.next
		}
	}

	// 从value的node中删除
	delete(node.keys, key)
	if len(node.keys) == 0 {
		// 没有key的值为node.value, 删除该节点
		if node.pre != nil {
			node.pre.next = node.next
		}
		if node.next != nil {
			node.next.pre = node.pre
		}
		if this.beg == node {
			this.beg = node.next
		}
		if this.end == node {
			this.end = node.next
		}
	}
}

func (this *AllOne) GetMaxKey() string {
	if this.beg == nil {
		return ""
	}
	for key := range this.beg.keys {
		return key
	}
	return ""
}

func (this *AllOne) GetMinKey() string {
	if this.end == nil {
		return ""
	}
	for key := range this.end.keys {
		return key
	}
	return ""
}
