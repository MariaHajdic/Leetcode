package main

func missingNumber(nums []int) int {
	next_expected, current_sum, expected_sum := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		current_sum += nums[i]
		expected_sum += next_expected
		next_expected++
	}
	return expected_sum + next_expected - current_sum
}

/* Slower solution */ /*
func missingNumber(nums []int) int {
	slices.Sort(nums)
	next_expected := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != next_expected {
			break
		}
		next_expected++
	}
	return next_expected
}*/
