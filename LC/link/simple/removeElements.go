package link

// RemoveElements 203. 移除链表元素
func RemoveElements(head *ListNode, val int) *ListNode {
	pre := &ListNode{Next: head}
	curr:=pre
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
			continue
		}
		curr = curr.Next
	}
	return pre.Next
}
