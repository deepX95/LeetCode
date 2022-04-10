package lru

type LRUCache struct {
	size, capacity int
	cache          map[int]*ListNode
	head, tail     *ListNode
}

type ListNode struct {
	key, val  int
	next, pre *ListNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		size:     0,
		head:     &ListNode{},
		tail:     &ListNode{},
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func (l *LRUCache) Get(key int) int {
	if _, hit := l.cache[key]; !hit {
		return -1
	}
	node := l.cache[key]
	l.moveToHead(node)
	return node.val
}

func (l *LRUCache) Put(key int, value int) {

}

func (l *LRUCache) addToHead(node *ListNode) {
	node.next = l.head.next
	l.head.next.pre = node
	l.head.next = node
	node.pre = l.head
}

func (l *LRUCache) moveToHead(node *ListNode) {
	l.removeNode(node)
	l.addToHead(node)
}

func (l *LRUCache) removeNode(node *ListNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
