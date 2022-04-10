package medium

type LinkNode struct{
	key, value int
	pre, next *LinkNode
}

// LRUCache 146. LRU 缓存
type LRUCache struct {
	m map[int]*LinkNode
	capacity int
	head, tail *LinkNode
}

func ConstructorLRUCache(capacity int) LRUCache {
	head := &LinkNode{-1, -1, nil, nil}
	tail := &LinkNode{-1, -1, nil, nil}
	head.next = tail
	tail.pre = head
	cache := LRUCache{make(map[int]*LinkNode), capacity, head, tail}
	return cache
}

func (c *LRUCache) AddNode(node *LinkNode) {
	node.pre = c.head
	node.next = c.head.next
	c.head.next = node
	node.next.pre = node
}

func (c *LRUCache) RemoveNode(node *LinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (c *LRUCache) MoveToHead(node *LinkNode) {
	c.RemoveNode(node)
	c.AddNode(node)
}

func (c *LRUCache) Get(key int) int {
	m := c.m
	if node, ok := m[key]; ok {
		c.MoveToHead(node)
		return node.value
	} else {
		return -1
	}
}

func (c *LRUCache) Put(key int, value int)  {
	m := c.m
	if node, ok := m[key]; ok {
		node.value = value
		c.MoveToHead(node)
	} else {
		n := &LinkNode{key, value, nil, nil}
		if len(m) >= c.capacity {
			delete(m, c.tail.pre.key)
			c.RemoveNode(c.tail.pre)
		}
		m[key] = n
		c.AddNode(n)
	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
