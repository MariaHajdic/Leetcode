package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	overflow := 0
	result := &ListNode{}
	ptr := result

	for l1 != nil || l2 != nil || overflow != 0 {
		sum := overflow

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		overflow = sum / 10
		sum %= 10
		ptr.Next = &ListNode{Val: sum}
		ptr = ptr.Next
	}
	return result.Next
}
