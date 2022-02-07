package link

// HasCycle 141 循环链表
func HasCycle(head *ListNode) bool {
	// 方法1
	//if head==nil{
	//	return false
	//}
	//fast,slow:=head.Next,head
	//for fast!=nil && fast.Next!=nil{
	//	fast=fast.Next.Next
	//	slow=slow.Next
	//	if fast==slow.Next{
	//		return true
	//	}
	//}
	//
	//return false

	// 方法2
	fast,slow:=head,head
	for fast!=nil && fast.Next!=nil{
		fast=fast.Next.Next
		slow=slow.Next
		if fast==slow {
			return true
		}
	}

	return false
}