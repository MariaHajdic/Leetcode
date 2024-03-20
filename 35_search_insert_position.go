package main

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

/*
func binarySearch(nums []int, idx, target int) int {
    if nums[idx] == target {
        return idx
    }
    if nums[idx] < target {
        binarySearch(nums, idx+idx/2, target)
    }
    binarySearch(nums, idx-idx/2, target)
    return 0
}
*/
