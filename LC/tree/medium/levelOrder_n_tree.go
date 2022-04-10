package medium

import "LeetCode/LC/tree"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// levelOrder 429. N 叉树的层序遍历
func levelOrderNTree(root *tree.Node) [][]int {
	if root == nil {
		return nil
	}

	var ret [][]int
	q := []*tree.Node{root}
	for q!=nil {
		var level []int
		tmp := q
		for _, node := range tmp {
			level = append(level, node.Val)
			q = append(q, node.Children...)
		}
		ret = append(ret, level)
	}
	return ret
}
