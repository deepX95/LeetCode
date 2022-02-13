package simple

import "LeetCode/LC/tree"

// InorderTraversal 94. 二叉树的中序遍历
func InorderTraversal(root *tree.TreeNode) (ret []int) {
	// 递归
	//if root==nil{
	//	return nil
	//}
	//ret=append(ret,InorderTraversal(root.Left)...)
	//ret=append(ret,root.Val)
	//ret=append(ret,InorderTraversal(root.Right)...)
	//return ret

	// 迭代
	var stack []*tree.TreeNode
	for root!=nil || len(stack)>0{
		for root!=nil{
			stack=append(stack,root)
			root=root.Left
		}
		root=stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		ret=append(ret,root.Val)
		root=root.Right
	}
	return
}