package main

func removeElement(nums []int, val int) int {
	start := 0
	for end := 0; end < len(nums); end++ {
		if nums[end] != val {
			nums[start] = nums[end]
			start++
		}
	}
	return start
}
