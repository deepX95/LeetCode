package test

import "math"

// MinStack 面试题 03.02. 栈的最小值
type MinStack struct {
	inStack []int
	min     int
}

/** initialize your data structure here. */
func ConstructorStack() MinStack {
	return MinStack{
		min: math.MaxInt64,
	}
}

func (s *MinStack) Push(x int) {
	// 对新增加的元素进行判断，是否小于当前最小值min
	if x <= s.min {
		// 当前添加的值为最小值时，则在栈中记录下上一个最小值
		// 这样在栈pop操作时，如果弹出当前最小值，就会将最小值更新为记录的上一个最小值
		s.inStack = append(s.inStack, s.min)
		// 最小值更新为当前值
		s.min = x
	}
	// 添加当前值到栈中
	s.inStack = append(s.inStack, x)
}

func (s *MinStack) Pop() {
	// 首先弹出最上方元素
	ele := s.inStack[len(s.inStack)-1]
	s.inStack = s.inStack[:len(s.inStack)-1]

	// 判断弹出的元素的值是否是当前最小值，如果是的话，就要更新当前最小值
	// 若当前弹出的元素不是当前最小值，那么仅仅弹出此元素就可以了，因为不影响当前最小值，不用更新
	if ele == s.min {
		// 因为在添加的元素为当前最小值时，都会在其添加之前通过逻辑添加该元素添加之前的最小值作为记录
		// 因此接着弹出栈的元素，这时弹出的元素为之前记录的上一个最小值，并将之更新为当前最小值
		// 举例：当前栈[MAX, -2, 0 ,-2, -3, 1, -3, -4]
		// (MAX, -2),(0),(-2, -3),(1),(-3, -4)栈中元素可以分为这几组
		// (MAX, -2),(-2, -3),(-3, -4)，这三组是push时，添加的元素为最小值，需要更新当前最小值，并且记录上一个最小值的
		// 当pop元素时，轮到这几组的时候，假如轮到(-3, -4)，pop出-4，此时最小值是-4，结果就是最小值被弹出了
		// 那么需要更新最小值，而上一个也就是在-4push添加进来之前最小值已经记录在-4的下方为-3，则将最小值更新为-3，
		// 因为在(-3, -4)这组中，-3属于为了记录上一个最小值而添加到栈中的，不属于真实要记录的数据，所以我们在更新最小值时，要连它-3一起弹出
		// (0),(1)，这俩组是在push时，添加的元素不是最小值，因此不需要更新当前最小值，更不需要记录上一个最小值
		s.min = s.inStack[len(s.inStack)-1]
		s.inStack = s.inStack[:len(s.inStack)-1]
	}
}

func (s *MinStack) Top() int {
	return s.inStack[len(s.inStack)-1]
}

func (s *MinStack) GetMin() int {
	return s.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
