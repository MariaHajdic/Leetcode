package main

func removeDuplicatesII(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	ptr1 := 2 // The first two are always allowed.

	for ptr2 := 2; ptr2 < len(nums); ptr2++ {
		// Covering a pair of allowed unique values.
		if nums[ptr2] != nums[ptr1-2] {
			nums[ptr1] = nums[ptr2]
			ptr1++ // The left side of the sliding window can now move.
		}
	}

	return ptr1
}
