package simple

import "LeetCode/LC/tree"

// PreorderTraversal 144. 二叉树的前序遍历
// fixme:题解 https://leetcode-cn.com/problems/binary-tree-preorder-traversal/solution/leetcodesuan-fa-xiu-lian-dong-hua-yan-shi-xbian-2/
func PreorderTraversal(root *tree.TreeNode) (ret []int){
	// 递归
	//if root==nil{
	//	return nil
	//}
	//ret=append(ret,root.Val)
	//ret=append(ret,preorderTraversal(root.Left)...)
	//ret=append(ret,preorderTraversal(root.Right)...)
	//return ret

	//var preorder func(*tree.TreeNode)
	//preorder = func(node *tree.TreeNode) {
	//	if node == nil {
	//		return
	//	}
	//	ret = append(ret, node.Val)
	//	preorder(node.Left)
	//	preorder(node.Right)
	//}
	//preorder(root)
	//return

	// 迭代
	stack:=[]*tree.TreeNode{}
	node:=root
	for node!=nil || len(stack)>0{
		for node!=nil{
			ret=append(ret,node.Val)
			stack=append(stack,node)
			node=node.Left
		}
		node=stack[len(stack)-1].Right
		stack=stack[:len(stack)-1]
	}
	return
}
