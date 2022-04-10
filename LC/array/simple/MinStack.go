package simple

// MinStack 面试题 03.02. 栈的最小值
type MinStack struct {
	inStack []int
	minStack[]     int
}

/** initialize your data structure here. */
func ConstructorStack() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(x int) {
	s.inStack=append(s.inStack,x)
	if len(s.minStack)!=0{
		tmp:=s.minStack[len(s.minStack)-1]
		if tmp<x{
			s.minStack=append(s.minStack,tmp)
		}else{
			s.minStack=append(s.minStack,x)
		}
	}else{
		s.minStack=append(s.minStack,x)
	}
}

func (s *MinStack) Pop() {
	if len(s.inStack)==0{
		return
	}
	s.inStack=s.inStack[:len(s.inStack)-1]
	s.minStack=s.minStack[:len(s.minStack)-1]
}

func (s *MinStack) Top() int {
	if len(s.inStack) != 0 {
		return s.inStack[len(s.inStack)-1]
	}
	return -1
}

func (s *MinStack) GetMin() int {
	return s.minStack[len(s.minStack)-1]
}



/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
