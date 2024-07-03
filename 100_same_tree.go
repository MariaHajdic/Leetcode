package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && q == nil
	}

	queue1, queue2 := []*TreeNode{p}, []*TreeNode{q}
	for len(queue1) > 0 && len(queue2) > 0 {
		curr1, curr2 := queue1[0], queue2[0] // the front nodes at the queues
		queue1, queue2 = queue1[1:], queue2[1:]

		if curr1.Val != curr2.Val {
			return false
		}

		if curr1.Left != nil || curr2.Left != nil {
			if curr1.Left == nil || curr2.Left == nil {
				return false
			}
			queue1 = append(queue1, curr1.Left)
			queue2 = append(queue2, curr2.Left)
		}

		if curr1.Right != nil || curr2.Right != nil {
			if curr1.Right == nil || curr2.Right == nil {
				return false
			}
			queue1 = append(queue1, curr1.Right)
			queue2 = append(queue2, curr2.Right)
		}
	}

	return len(queue1) == 0 && len(queue2) == 0
}
