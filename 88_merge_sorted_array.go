package main

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := range n {
		nums1[m] = nums2[i]
		m++
	}
	sort.Ints(nums1)
}

/* This one is slower.

func merge(nums1 []int, m int, nums2 []int, n int)  {
    ptr1, ptr2, p := m-1, n-1, m+n-1

    for ; ptr1 >= 0 && ptr2 >= 0; p-- {
        if nums1[ptr1] > nums2[ptr2] {
            nums1[p] = nums1[ptr1]
            ptr1--
        } else {
            nums1[p] = nums2[ptr2]
            ptr2--
        }
    }

    for ; ptr2 >= 0; ptr2-- {
        nums1[p] = nums2[ptr2]
        p--
    }
}
*/
