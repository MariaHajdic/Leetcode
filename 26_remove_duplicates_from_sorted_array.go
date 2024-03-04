package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	ptr1 := 1

	for ptr2 := 1; ptr2 < len(nums); ptr2++ {
		if nums[ptr2] != nums[ptr1-1] {
			nums[ptr1] = nums[ptr2]
			ptr1++
		}
	}
	return ptr1
}
