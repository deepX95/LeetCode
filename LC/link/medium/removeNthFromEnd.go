package medium

import "LeetCode/LC/link"

// RemoveNthFromEnd   19. 删除链表的倒数第 N 个结点
func RemoveNthFromEnd(head *link.ListNode, n int) *link.ListNode {
	// 根据长度删除
	//lth := 0
	//for tmp:=head;tmp != nil;tmp=tmp.Next {
	//	lth++
	//}
	//
	//dummy := &link.ListNode{Next: head}
	//cur := dummy
	//for i:=0; i<lth-n; i++{
	//	cur = cur.Next
	//}
	//cur.Next = cur.Next.Next
	//return dummy.Next

	// 双指针
	dummy:=&link.ListNode{Next: head}
	first,second:=head,dummy
	for i:=0;i<n;i++{
		first=first.Next
	}
	for ;first!=nil;first=first.Next{
		second=second.Next
	}
	second.Next=second.Next.Next
	return dummy.Next

}
