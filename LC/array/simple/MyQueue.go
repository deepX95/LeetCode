package simple

// MyQueue 用栈实现队列
type MyQueue struct {
	in, out []int
}

func ConstructorQueue() MyQueue {
	return MyQueue{}
}

// Push 将元素 x 推到队列的末尾
func (q *MyQueue) Push(x int) {
	q.in = append(q.in, x)
}

func (q *MyQueue) PushOut() {
	for len(q.in) > 0 {
		q.out = append(q.out, q.in[len(q.in)-1])
		q.in = q.in[:len(q.in)-1]
	}
}

// Pop 从队列的开头移除并返回元素
func (q *MyQueue) Pop() int {
	if len(q.out) == 0 {
		q.PushOut()
	}
	x := q.out[len(q.out)-1]
	q.out = q.out[:len(q.out)-1]
	return x
}

// Peek 返回队列开头的元素
func (q *MyQueue) Peek() int {
	if len(q.out) == 0 {
		q.PushOut()
	}
	return q.out[len(q.out)-1]
}

// Empty() 判空
func (q *MyQueue) Empty() bool {
	return len(q.in) == 0 && len(q.out) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
