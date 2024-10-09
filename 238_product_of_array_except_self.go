package main

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums)) // left products
	curr := 1
	for i := 0; i < len(nums); i++ {
		res[i] = curr
		curr *= nums[i]
	}

	curr = 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= curr
		curr *= nums[i]
	}

	return res
}
