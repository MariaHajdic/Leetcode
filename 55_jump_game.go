package main

func canJump(nums []int) bool {
	target := len(nums) - 1
	for i := target; i >= 0; i-- {
		if nums[i]+i >= target {
			target = i
		}
	}
	return target == 0
}
