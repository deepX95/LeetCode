package simple

import "LeetCode/LC/link"

// GetKthFromEnd 剑指 Offer 22. 链表中倒数第k个节点
func GetKthFromEnd(head *link.ListNode, k int) *link.ListNode {
	// 顺序查找
	//n:=0
	//for node:=head;node!=nil;node=node.Next{
	//	n++
	//}
	//curr:=head
	//for ;n>k;n--{
	//	curr=curr.Next
	//}
	//return curr

	// 快慢指针
	slow,fast:=head,head
	for fast!=nil && k>0{
		fast=fast.Next
		k--
	}
	for fast!=nil{
		slow=slow.Next
		fast=fast.Next
	}
	return slow
}
