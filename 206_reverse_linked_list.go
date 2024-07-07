package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	values := []int{}
	current_node := head
	for current_node.Next != nil {
		values = append(values, current_node.Val)
		current_node = current_node.Next
	}
	values = append(values, current_node.Val) // covers the last one

	current_node = head
	for i := len(values) - 1; i >= 0; i-- {
		current_node.Val = values[i]
		current_node = current_node.Next
	}

	return head
}
