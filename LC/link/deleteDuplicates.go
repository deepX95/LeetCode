package link

// DeleteDuplicates 83. 删除排序链表中的重复元素
func DeleteDuplicates(head *ListNode) *ListNode {
	// 1.题解官方答案
	cur := head
	for cur!=nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head

	// 我的答案
	curr:=head
	for curr!=nil && curr.Next!=nil{
		fast:=curr.Next
		for fast!=nil {
			if curr.Val ==fast.Val {
				curr.Next =fast.Next
			}
			fast=fast.Next
		}
		curr=curr.Next

	}
	return head
}