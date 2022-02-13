package simple

import "LeetCode/LC/tree"

// postorderTraversal 145. 二叉树的后序遍历
// fixme:题解：https://leetcode-cn.com/problems/binary-tree-postorder-traversal/solution/er-cha-shu-de-hou-xu-bian-li-by-leetcode-solution/
func postorderTraversal(root *tree.TreeNode) (ret []int) {
	// 递归
	//if root==nil{
	//	return nil
	//}
	//ret=append(ret,postorderTraversal(root.Left)...)
	//ret=append(ret,postorderTraversal(root.Right)...)
	//ret=append(ret,root.Val)
	//return ret

	// 迭代
	var stack []*tree.TreeNode
	var prev *tree.TreeNode
	for root!=nil ||len(stack)>0{
		// 左节点压栈
		for root!=nil{
			stack=append(stack,root)
			root=root.Left
		}

		// 取栈顶元素
		root=stack[len(stack)-1]
		// 出栈
		stack=stack[:len(stack)-1]
		if root.Right==nil || root.Right==prev{
			ret=append(ret,root.Val)
			prev=root
			root=nil
		}else{
			stack=append(stack,root)
			root=root.Right
		}
	}
	return
}