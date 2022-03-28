package medium

//type MyLinkedList struct {
//	root *Element
//	len  int
//}
//
//type Element struct {
//	next, pre *Element
//	val       int
//}
//
//func Constructor() MyLinkedList {
//	link := MyLinkedList{
//		root: &Element{},
//	}
//	link.root.next = link.root
//	link.root.pre = link.root
//	link.len = 0
//
//	return link
//}
//
//// Get 获取链表中第 index 个节点的值。如果索引无效，则返回-1
//func (l *MyLinkedList) Get(index int) int {
//	if l.len == 0 || l.len <= index {
//		return -1
//	}
//
//	cnt, cur := 0, l.root
//	for cnt < index {
//		cur = cur.next
//		cnt++
//	}
//	if cnt == index {
//		return cur.val
//	}
//
//	return -1
//}
//
//// AddAtHead 在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点
//func (l *MyLinkedList) AddAtHead(val int) {
//	ele := &Element{
//		val: val,
//	}
//
//	if l.len == 0 {
//		l.root = ele
//		l.root.next = l.root
//		l.root.pre = l.root
//		l.len++
//		return
//	}
//
//	l.len++
//	ele.pre = l.root.pre
//	ele.next = l.root
//	l.root.pre.next = ele
//	l.root.pre = ele
//
//	l.root = ele
//}
//
//// AddAtTail 将值为 val 的节点追加到链表的最后一个元素
//func (l *MyLinkedList) AddAtTail(val int) {
//	ele := &Element{
//		val: val,
//	}
//
//	if l.len == 0 {
//		l.root = ele
//		l.root.next = ele
//		l.root.pre = ele
//		l.len++
//		return
//	}
//
//	l.len++
//
//	ele.pre = l.root.pre
//	ele.next = l.root
//
//	l.root.pre.next = ele
//	l.root.pre = ele
//
//}
//
//func (l *MyLinkedList) AddAtIndex(index int, val int) {
//	if l.len == index {
//		l.AddAtTail(val)
//		return
//	} else if index <= 0 {
//		l.AddAtHead(val)
//		return
//	} else if l.len < index {
//		return
//	}
//
//	ele := &Element{
//		val: val,
//	}
//	l.len++
//	cnt, cur := 0, l.root
//	for cnt < index {
//		cnt++
//		cur = cur.next
//	}
//	ele.next = cur
//	ele.pre = cur.pre
//
//	cur.pre.next = ele
//	cur.pre = ele
//}
//
//func (l *MyLinkedList) DeleteAtIndex(index int) {
//	if l.len == 0 || l.len-1 < index || index < 0 {
//		return
//	}
//
//	if l.len-1 == index { // 尾
//		// l.root 根节点
//		// l.root.pre  尾节点
//		// l.root.pre.pre 尾节点的上一个节点
//
//		l.root.pre.pre.next = l.root
//		l.root.pre = l.root.pre.pre
//		l.len--
//		return
//	}
//
//	if index == 0 { // 头
//		// l.root 根节点
//		//l.root.pre  尾节点
//		l.root.pre.next = l.root.next
//		l.root.next.pre = l.root.pre
//		l.root = l.root.next
//		l.len--
//		return
//	}
//	cnt, cur := 0, l.root
//	for cnt < index {
//		cur = cur.next
//		cnt++
//	}
//
//	cur.pre.next = cur.next
//	cur.next.pre = cur.pre
//	l.len--
//	return
//}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

// 单链表实现
//type ListNode struct {
//	val  int
//	next *ListNode
//}
//
//type MyLinkedList struct {
//	size int
//	head *ListNode // 哨兵节点,统一头节点的行为
//}
//
//func Constructor() MyLinkedList {
//	return MyLinkedList{
//		size: 0,
//		head: &ListNode{},
//	}
//}
//
//func (l *MyLinkedList) Get(index int) int {
//	if index < 0 || index >= l.size {
//		return -1
//	}
//	cur := l.head
//	for i := 0; i < index+1; i++ {
//		cur = cur.next
//	}
//	return cur.val
//}
//
//func (l *MyLinkedList) AddAtHead(val int) {
//	l.AddAtIndex(0, val)
//}
//
//func (l *MyLinkedList) AddAtTail(val int) {
//	l.AddAtIndex(l.size, val)
//}
//
//func (l *MyLinkedList) AddAtIndex(index int, val int) {
//	if index < 0 || index > l.size {
//		return
//	}
//	prev := l.head
//	for i := 0; i < index; i++ {
//		prev = prev.next
//	}
//	node := &ListNode{val: val}
//	node.next = prev.next
//	prev.next = node
//	l.size++
//}
//
//func (l *MyLinkedList) DeleteAtIndex(index int) {
//	if index < 0 || index >= l.size {
//		return
//	}
//	prev := l.head
//	for i := 0; i < index; i++ {
//		prev = prev.next
//	}
//	prev.next = prev.next.next
//	l.size--
//}

// 双链表实现
type MyLinkedList struct {
	size int
	head *ListNode
	tail *ListNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		size: 0,
		head: &ListNode{},
		tail: &ListNode{},
	}
}

type ListNode struct {
	val  int
	next *ListNode
	prev *ListNode
}

func (l *MyLinkedList) Get(index int) int {

}

func (l *MyLinkedList) AddAtHead(val int) {

}

func (l *MyLinkedList) AddAtTail(val int) {

}

func (l *MyLinkedList) AddAtIndex(index int, val int) {

}

func (l *MyLinkedList) DeleteAtIndex(index int) {

}
