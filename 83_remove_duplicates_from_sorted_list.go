package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	current_node := head

	for current_node != nil && current_node.Next != nil {
		if current_node.Val == current_node.Next.Val {
			current_node.Next = current_node.Next.Next
			continue
		}
		current_node = current_node.Next
	}

	return head
}
