package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		left, right := queue[0], queue[1]
		queue = queue[2:]

		if left != nil {
			if left.Left == nil && left.Right == nil {
				sum += left.Val
			} else {
				queue = append(queue, left.Left, left.Right)
			}
		}
		if right != nil {
			queue = append(queue, right.Left, right.Right)
		}
	}

	return sum
}
