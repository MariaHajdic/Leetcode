package main

import "slices"

func containsDuplicate(nums []int) bool {
	slices.Sort(nums)
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == prev {
			return true
		}
		prev = nums[i]
	}
	return false
}
