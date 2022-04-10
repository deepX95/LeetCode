package difficult

import "LeetCode/LC/link"

// reverseKGroup 25. K 个一组翻转链表
// fixme:多做几次
func reverseKGroup(head *link.ListNode, k int) *link.ListNode {
	hair := &link.ListNode{Next: head}
	pre := hair
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		next := tail.Next
		head, tail = reverse(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

func reverse(head, tail *link.ListNode) (*link.ListNode, *link.ListNode) {
	pre := tail.Next
	p := head
	for pre != tail {
		next := p.Next
		p.Next = pre
		pre = p
		p = next
	}
	return tail, head
}
