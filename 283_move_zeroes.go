package main

func moveZeroes(nums []int) {
	zero, number := 0, 0
	for zero < len(nums) && number < len(nums) {
		if nums[zero] != 0 {
			zero++
			continue
		}
		if nums[number] == 0 {
			number++
			continue
		}
		if zero > number {
			number++
			continue
		}
		nums[zero], nums[number] = nums[number], 0
		zero++
		number++
	}
}
