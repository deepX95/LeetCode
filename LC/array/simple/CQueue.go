package simple

// 剑指 Offer 09. 用两个栈实现队列
type CQueue struct {
	inStack, outStack []int
}

func ConstructorCQueue() CQueue {
	return CQueue{}
}

func (q *CQueue) PushOut() {
	for len(q.inStack) > 0 {
		q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
		q.inStack = q.inStack[:len(q.inStack)-1]
	}
}

func (q *CQueue) AppendTail(value int) {
	q.inStack = append(q.inStack, value)
}

func (q *CQueue) DeleteHead() int {
	if len(q.outStack) == 0 {
		if len(q.inStack)==0{
			return -1
		}
		q.PushOut()
	}

	del := q.outStack[len(q.inStack)-1]
	q.outStack = q.outStack[:len(q.inStack)-1]
	return del
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
