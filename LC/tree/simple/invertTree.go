package simple

import "LeetCode/LC/tree"

// invertTree  226. 翻转二叉树
// 使用迭代
func invertTree(root *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}
