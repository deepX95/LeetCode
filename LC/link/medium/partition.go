package medium

import "LeetCode/LC/link"

// Partition 86. 分隔链表
func Partition(head *link.ListNode, x int) *link.ListNode {
	small,large:=&link.ListNode{},&link.ListNode{}
	smallH,largeH:=small,large
	for head!=nil{
		if head.Val<x{
			small.Next=head
			small=small.Next
		}else{
			large.Next=head
			large=large.Next
		}
		head=head.Next
	}
	large.Next=nil
	small.Next=largeH.Next
	return smallH.Next
}
