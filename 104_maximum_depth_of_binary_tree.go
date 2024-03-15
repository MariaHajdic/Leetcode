package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	/*
	   2^0: 3
	   2^1: 9, 20
	   2^2: null, null, 15, 7
	   len(root) - sum
	*/
	// depth, length := 0, len(root)
	// for length != 0 {
	//     depth++
	//     length >> 1
	// }
	// return depth
	if root == nil {
		return 0
	}
	return 1 + int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))))
}
