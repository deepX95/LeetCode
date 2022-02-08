package medium

import "LeetCode/LC/link"

// MergeInBetween 1669. 合并两个链表
func MergeInBetween(list1 *link.ListNode, a int, b int, list2 *link.ListNode) *link.ListNode {
	// 1.遍历
	//head2:=list2  // list2 头结点
	//tail:=head2   // list2 尾结点
	//for tail.Next!=nil{
	//	tail=tail.Next
	//}
	//
	//// 得前一个节点
	//n:=1
	//currStart :=list1
	//for n<a{
	//	currStart=currStart.Next
	//	n++
	//	if n==a{
	//		break
	//	}
	//}
	//
	//// 得当前节点
	//n=0
	//currEnd :=list1
	//for n<b{
	//	currEnd=currEnd.Next
	//	n++
	//	if n==b{
	//		break
	//	}
	//}
	//currStart.Next=head2
	//tail.Next=currEnd.Next
	//
	//return list1


	// 2.遍历优化
	//temp1:=list1
	//temp2:=list1
	//for i:=0;i<a-1;i++{
	//	temp1=temp1.Next
	//}
	//for j:=0;j<b;j++{
	//	temp2=temp2.Next
	//}
	//temp1.Next=list2
	//temp3:=list2
	//for temp3.Next!=nil{
	//	temp3=temp3.Next
	//}
	//temp3.Next=temp2.Next
	//return list1

	// 3.快慢指针
	rear2 := list2
	//rear2指向list2的尾节点
	for rear2.Next != nil {
		rear2 = rear2.Next
	}
	//哑节点
	dummyHead := &link.ListNode{
		Next: list1,
	}
	pre1, pre2 := dummyHead, dummyHead
	step := 0
	//pre1指向第a-1个节点，pre2指向第b个节点
	for ptr := dummyHead; ptr != nil && step <= b; ptr = ptr.Next {
		if step < a {
			pre1 = pre1.Next
		}
		pre2 = pre2.Next
		step++
	}
	//连接
	pre1.Next = list2
	rear2.Next = pre2.Next
	return dummyHead.Next
}
