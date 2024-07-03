package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2 // be that nil or a tree
	}
	if root2 == nil {
		return root1
	}

	queue1, queue2 := []*TreeNode{root1}, []*TreeNode{root2}
	for len(queue1) > 0 && len(queue2) > 0 {
		curr1, curr2 := queue1[0], queue2[0]
		queue1, queue2 = queue1[1:], queue2[1:]
		curr1.Val += curr2.Val

		if curr1.Left != nil && curr2.Left != nil { // the only reason to go deeper
			queue1 = append(queue1, curr1.Left)
			queue2 = append(queue2, curr2.Left)
		} else if curr1.Left == nil && curr2.Left != nil {
			curr1.Left = curr2.Left
		}
		if curr1.Right != nil && curr2.Right != nil { // same for rights
			queue1 = append(queue1, curr1.Right)
			queue2 = append(queue2, curr2.Right)
		} else if curr1.Right == nil && curr2.Right != nil {
			curr1.Right = curr2.Right
		}
	}

	return root1
}
