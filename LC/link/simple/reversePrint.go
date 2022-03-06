package simple

import "LeetCode/LC/link"

func reversePrint(head *link.ListNode) []int {
	// 1.栈
	//var stack []int
	//for head != nil {
	//	stack = append(stack, head.Val)
	//	head = head.Next
	//}
	//l, r := 0, len(stack)-1
	//for l < r {
	//	stack[l], stack[r] = stack[r], stack[l]
	//	l++
	//	r--
	//}
	//return stack

	// 递归
	var res []int
	var recur func(*link.ListNode)
	recur = func(head *link.ListNode) {
		if head == nil {
			return
		}
		recur(head.Next)
		res = append(res, head.Val)
	}
	recur(head)
	return res
}
