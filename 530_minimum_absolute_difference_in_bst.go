package main

import "slices"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getMinimumDifference(root *TreeNode) int {
	vals := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		vals = append(vals, current.Val)

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	slices.Sort(vals)
	minDiff := vals[1] - vals[0]

	for i := 2; i < len(vals); i++ {
		diff := vals[i] - vals[i-1]
		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}

// wrong solution
// func getMinimumDifference(root *TreeNode) int {
//     var minDiff int
//     if root.Left != nil {
//         minDiff = root.Val - root.Left.Val
//     } else {
//         minDiff = root.Right.Val - root.Val
//     }

//     queue := []*TreeNode{root}
//     for len(queue) > 0 {
//         current := queue[0]
//         queue = queue[1:]

//         if current.Left != nil {
//             diff := current.Val - current.Left.Val
//             if diff < minDiff {
//                 minDiff = diff
//             }
//             queue = append(queue, current.Left)
//         }
//         if current.Right != nil {
//             diff := current.Right.Val - current.Val
//             if diff < minDiff {
//                 minDiff = diff
//             }
//             queue = append(queue, current.Right)
//         }

//         if minDiff == 0 {
//             return 0
//         }
//     }

//     return minDiff
// }
/*
        263
        /  \
      104  701
     /  \  /  \
  nil 227  nil 911
*/
