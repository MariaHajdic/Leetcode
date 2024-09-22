package main

func search(nums []int, target int) int {
	for left, right := 0, len(nums)-1; left <= right; {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
			continue
		}
		left = mid + 1
	}

	return -1
}
