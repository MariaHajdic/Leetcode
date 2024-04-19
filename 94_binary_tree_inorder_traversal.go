package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	currentNode := root

	for currentNode != nil {
		if currentNode.Left == nil {
			result = append(result, currentNode.Val)
			currentNode = currentNode.Right
			continue
		}
		predecessor := currentNode.Left
		for predecessor.Right != nil && predecessor.Right != currentNode {
			predecessor = predecessor.Right
		}

		if predecessor.Right == nil {
			predecessor.Right = currentNode
			currentNode = currentNode.Left
			continue
		}
		predecessor.Right = nil
		result = append(result, currentNode.Val)
		currentNode = currentNode.Right
	}

	return result
}
