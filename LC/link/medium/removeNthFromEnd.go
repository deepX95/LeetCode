package medium

import "LeetCode/LC/link"

// RemoveNthFromEnd   19. 删除链表的倒数第 N 个结点
func RemoveNthFromEnd(head *link.ListNode, n int) *link.ListNode {
	// 根据长度删除
	k := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		k++
	}

	// 删除头结点
	if k == n {
		return head.Next
	}

	cur = head
	for ; k-n-1 > 0; k-- {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return head

	// 双指针

}
