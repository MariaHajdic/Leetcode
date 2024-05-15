package main

import "slices"

func singleNumberii(nums []int) int {
	slices.Sort(nums)
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == prev {
			i++
			continue
		}
		if i == 1 && nums[i] != prev {
			return prev
		}
		prev = nums[i]
		if i < len(nums)-1 && nums[i+1] != prev {
			break
		}
	}
	return prev
}
