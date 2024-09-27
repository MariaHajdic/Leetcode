package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := [][2]interface{}{{root, 1}}

	for len(queue) > 0 {
		current := queue[0][0].(*TreeNode)
		depth := queue[0][1].(int)
		queue = queue[1:]

		if current.Left == nil && current.Right == nil {
			return depth
		}

		if current.Left != nil {
			queue = append(queue, [2]interface{}{current.Left, depth + 1})
		}
		if current.Right != nil {
			queue = append(queue, [2]interface{}{current.Right, depth + 1})
		}
	}

	return 0
}
