package main

func sortColors(nums []int) {
	left, right := 0, len(nums)-1

	for i := left; i <= right; {
		if nums[i] == 1 {
			i++
			continue
		}
		if nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right--
			continue
		}
		if nums[i] == 0 {
			nums[left], nums[i] = nums[i], nums[left]
			left++
			i++
		}
	}
}
