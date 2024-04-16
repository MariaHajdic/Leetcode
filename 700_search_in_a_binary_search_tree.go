package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		if stack[len(stack)-1].Val == val {
			return stack[len(stack)-1]
		}
		if stack[len(stack)-1].Left != nil {
			stack = append(stack, stack[len(stack)-1].Left)
			stack[len(stack)-2].Left = nil
			continue
		}
		if stack[len(stack)-1].Right != nil {
			stack = append(stack, stack[len(stack)-1].Right)
			stack[len(stack)-2].Right = nil
			continue
		}
		stack = stack[:len(stack)-1]
	}
	return nil
}
