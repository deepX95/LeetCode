package main

import (
	"LeetCode/LC/link"
)

func main() {
	//LC.TwoSum([]int{3,2,4},6)
	//LC.ReverseString([]string{"H"})
	//fmt.Println(LC.PlusOne([]int{1,9,9}))

	//nodes := &link.ListNode{
	//	Val: 1,
	//	Next: &link.ListNode{
	//		Val: 2,
	//		Next: &link.ListNode{
	//			Val: 3,
	//			Next: &link.ListNode{
	//				Val: 4,
	//				Next: &link.ListNode{
	//					Val:  5,
	//					Next: nil,
	//				},
	//			},
	//		},
	//	},
	//}
	//
	////ret:=LC.ReverseList(nodes)
	//link.DeleteNode(nodes,3)
	comNode:=&link.ListNode{
				Val:4,
				Next: &link.ListNode{
					Val: 5,
					Next: &link.ListNode{
						Val:  4,
						Next: nil,
					},
				},
			}


	headA := &link.ListNode{
		Val: 2,
		Next: &link.ListNode{
			Val: 2,
			Next: comNode,
		},
	}
	headB :=&link.ListNode{
		Val: 2,
		Next: &link.ListNode{
			Val: 2,
			Next: comNode,
		},
	}

	link.GetIntersectionNode(headA,headB)
}
