package simple

import "LeetCode/LC/link"

// isPalindrome 234. 回文链表
func isPalindrome(head *link.ListNode) bool {
	// 1.转成数组
	var val []int
	for head!=nil{
		val=append(val,head.Val)
		head=head.Next
	}
	n := len(val)
	for i, v := range val[:n/2] {
		if v != val[n-1-i] {
			return false
		}
	}
	return true


}
