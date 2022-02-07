package link

// DeleteNode 删除节点
func DeleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}

	curr := head
	for curr != nil && curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
			return head
		}
		curr = curr.Next
	}

	return head
}
