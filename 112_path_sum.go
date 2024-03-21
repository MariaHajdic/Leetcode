package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		current := stack[0]
		stack = stack[1:]

		if current.Left == nil && current.Right == nil && current.Val == targetSum {
			return true
		}

		if current.Left != nil {
			stack = append(stack, current.Left)
			current.Left.Val += current.Val
		}
		if current.Right != nil {
			stack = append(stack, current.Right)
			current.Right.Val += current.Val
		}
	}

	return false
}
