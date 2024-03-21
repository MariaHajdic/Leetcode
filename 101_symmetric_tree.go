package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		left, right := queue[0], queue[1]
		queue = queue[2:]

		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}

		if left.Val != right.Val {
			return false
		}

		queue = append(queue, left.Left, right.Right)
		queue = append(queue, left.Right, right.Left)
	}

	return true
}

/* Recursive solution */
/*func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }

    return isMirror(root.Left, root.Right)
}

func isMirror(left, right *TreeNode) bool {
    if left == nil && right == nil {
        return true
    }
    if left == nil || right == nil {
        return false
    }
    return left.Val == right.Val && isMirror(left.Right, right.Left) && isMirror(left.Left, right.Right)
}*/
