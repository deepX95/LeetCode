package medium

import "LeetCode/LC/link"

func detectCycle(head *link.ListNode) *link.ListNode {
	// 哈希表
	//visit:=map[*link.ListNode]bool{}
	//for head!=nil{
	//	if ok:=visit[head];ok{
	//		return head
	//	}
	//	visit[head]=true
	//	head=head.Next
	//}
	//return nil

	// 快慢指针
	slow,fast:=head,head
	for fast!=nil && fast.Next!=nil{
		fast=fast.Next.Next
		slow=slow.Next

		if fast==slow{
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil

}
