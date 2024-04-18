package main

func pivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	leftSum := 0
	for i, num := range nums {
		if float64(leftSum) == float64(total-num)/float64(2) {
			return i
		}
		leftSum += num
	}

	return -1
}
