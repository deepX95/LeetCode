package simple

import "LeetCode/LC/link"

// DeleteNode2 237. 删除链表中的节点
func DeleteNode2(node *link.ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}