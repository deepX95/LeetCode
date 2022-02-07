package simple

import "LeetCode/LC/link"

// DeleteNode 删除节点
func DeleteNode(head *link.ListNode, val int) *link.ListNode {
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
