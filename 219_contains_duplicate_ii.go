package main

import "math"

func containsNearbyDuplicate(nums []int, k int) bool {
	duplicates := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		val, ok := duplicates[nums[i]]
		if ok {
			if int(math.Abs(float64(val-i))) <= k {
				return true
			}
		}
		duplicates[nums[i]] = i
	}
	return false
}
