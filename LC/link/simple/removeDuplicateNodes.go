package simple

import "LeetCode/LC/link"

// RemoveDuplicateNodes 面试题 02.01. 移除重复节点
func RemoveDuplicateNodes(head *link.ListNode) *link.ListNode {
	// 1.哈希表
	//if head==nil{
	//	return nil
	//}
	//
	//tags:=map[int]bool{
	//	head.Val:true,
	//}
	//cur:=head
	//for cur.Next!=nil{
	//	if _,ok:=tags[cur.Next.Val];!ok{
	//		tags[cur.Next.Val]=true
	//		cur=cur.Next
	//	}else{
	//		cur.Next=cur.Next.Next
	//	}
	//}
	//return head

	// 2.双重循环
	oc:=head
	for oc!=nil{
		ic:=oc
		for ic.Next!=nil{
			if ic.Next.Val==oc.Val{
				ic.Next=ic.Next.Next
				continue
			}
			ic=ic.Next
		}
		oc=oc.Next
	}
	return head
}
