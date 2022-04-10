package medium

import "LeetCode/LC/tree"

// levelOrderBottom 107. 二叉树的层序遍历 II
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *tree.TreeNode) [][]int {
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
	//反转二维数组
	for i := 0; i < len(ret) / 2; i++ {
		ret[i], ret[len(ret) - 1 - i] = ret[len(ret) - 1 - i], ret[i]
	}
	return ret
}