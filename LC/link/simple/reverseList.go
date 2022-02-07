package simple

import "LeetCode/LC/link"

func ReverseList(head *link.ListNode) *link.ListNode {
	var prev *link.ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
