package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rangeSumBST(root *TreeNode, low int, high int) int {
	stack := []*TreeNode{root}
	sum := 0

	for len(stack) > 0 {
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
		if stack[len(stack)-1].Val >= low && stack[len(stack)-1].Val <= high {
			sum += stack[len(stack)-1].Val
		}
		stack = stack[:len(stack)-1]
	}
	return sum
}
