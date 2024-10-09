package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil

	left := sortList(head)
	right := sortList(mid)

	return mergeLists(left, right)
}

func mergeLists(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	} else if l2 != nil {
		current.Next = l2
	}

	return head.Next
}

/* Classic merge sort.
func mergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    mid := len(arr) / 2
    left := mergeSort(arr[:mid])
    right := mergeSort(arr[mid:])

    return merge(left, right)
}

func merge(left, right []int) []int {
    result := make([]int, len(left)+len(right))

    i, j := 0, 0
    for i < len(left) && j < len(right) {
        if left[i] <= right[j] {
            result = append(result, left[i])
            i++
            continue
        }
        result = append(result, right[j])
        j++
    }
    if i < len(left) {
        result = append(result, left[i:])
    } else if j < len(right) {
        result = append(result, right[j:])
    }

    return result
}*/
