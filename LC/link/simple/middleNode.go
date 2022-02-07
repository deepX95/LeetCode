package simple

import "LeetCode/LC/link"

// MiddleNode 876. 链表的中间结点
func MiddleNode(head *link.ListNode) *link.ListNode {
	slow,fast:=head,head
	for fast!=nil && fast.Next!=nil{
		slow=slow.Next
		fast=fast.Next.Next
	}
	return slow
}