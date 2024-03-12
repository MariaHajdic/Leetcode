package main

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) < 1 {
		return nil
	}
	low, prev := nums[0], nums[0]
	result := []string{}

	for i := 1; i < len(nums); i++ {
		if nums[i] != prev+1 {
			result = append(result, appendRange(low, prev))
			low = nums[i]
		}
		prev = nums[i]
	}

	return append(result, appendRange(low, prev))
}

func appendRange(low, high int) string {
	s := fmt.Sprintf("%d", low)
	if low != high {
		s += fmt.Sprintf("->%d", high)
	}
	return s
}
