package simple

import "LeetCode/LC/link"

// RemoveElements 203. 移除链表元素
func RemoveElements(head *link.ListNode, val int) *link.ListNode {
	pre := &link.ListNode{Next: head}
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
