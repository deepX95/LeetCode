package simple

import "LeetCode/LC/tree"

// Preorder 589. N 叉树的前序遍历
func Preorder(root *tree.Node) (ret []int) {
	// 迭代
	if root==nil{
		return
	}
	stack:=[]*tree.Node{root}
	for len(stack)>0{
		node:=stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		ret=append(ret,node.Val)
		if len(node.Children)!=0{
			for i:=len(node.Children)-1;i>=0;i--{
				stack = append(stack, node.Children[i])
			}
		}
	}

	return

	// 递归
	//if root == nil{
	//	return nil
	//}
	//var ans []int
	//ans = append(ans,root.Val)
	//for _,node := range root.Children {
	//	ans = append(ans,Preorder(node)...)
	//}
	//return ans
}