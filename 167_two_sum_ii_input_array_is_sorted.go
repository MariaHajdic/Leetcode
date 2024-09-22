package main

func twoSumii(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for numbers[left]+numbers[right] != target {
		if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
	}
	return []int{left + 1, right + 1}
}
