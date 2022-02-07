package link

// MergeTwoLists 21. 合并两个有序链表
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head:=&ListNode{}
	cursor :=head
	for l1!=nil && l2!=nil{
		if l1.Val<=l2.Val{
			cursor.Next=l1
			l1=l1.Next
		}else{
			cursor .Next=l2
			l2=l2.Next
		}
		cursor =cursor.Next
	}

	if l1!=nil{
		cursor.Next=l1
	}else if l2!=nil{
		cursor.Next=l2
	}

	return head.Next
}