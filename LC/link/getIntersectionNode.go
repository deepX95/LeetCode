package link

// GetIntersectionNode 160.相交链表
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	// 双指针
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa

	// 哈希表
	//tags:=make(map[*ListNode]bool,0)
	//for headA!=headB{
	//	if headA!=nil {
	//		if tags[headA] {
	//			return headA
	//		} else {
	//			tags[headA]=true
	//			headA = headA.Next
	//		}
	//	}
	//
	//	if headB!=nil {
	//		if tags[headB] {
	//			return headB
	//		} else {
	//			tags[headB]=true
	//			headB = headB.Next
	//		}
	//	}
	//}
	//if headB!=nil &&headA!=nil{
	//	return headB
	//}
	//return nil

	// 哈希表
	//tags:=make(map[*ListNode]bool,0)
	//for headA!=nil{
	//	tags[headA]=true
	//	headA=headA.Next
	//}
	//for headB!=nil{
	//	if tags[headB]{
	//		return headB
	//	}
	//	headB=headB.Next
	//}
	//return nil
}