package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodes := map[*ListNode]bool{}

	for headA != nil {
		nodes[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if nodes[headB] == true {
			return headB
		}
		headB = headB.Next
	}

	return nil
}

/*
Extra space of len(A), but O(m+n)
*/
