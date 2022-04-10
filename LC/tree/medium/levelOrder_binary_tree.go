package medium

import "LeetCode/LC/tree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// levelOrder 102. 二叉树的层序遍历
func levelOrder(root *tree.TreeNode) [][]int {
	// bfs
	if root==nil{
		return nil
	}
	var ret [][]int
	q := []*tree.TreeNode{root}
	for len(q)!=0{
		var tmp []int
		nodes:=q
		q=nil
		for _,node:=range nodes{
			tmp=append(tmp,node.Val)
			if node.Left!=nil{
				q=append(q,node.Left)
			}
			if node.Right!=nil{
				q=append(q,node.Right)
			}
		}
		ret=append(ret,tmp)
	}
	return ret

	// fixme:DFS
}
