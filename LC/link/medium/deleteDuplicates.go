package medium

import "LeetCode/LC/link"

// DeleteDuplicates 82. 删除排序链表中的重复元素 II
// fixme:需要再看几遍
func DeleteDuplicates(head *link.ListNode) *link.ListNode {
	dummy := &link.ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			val := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == val {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
