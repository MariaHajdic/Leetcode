package main

func findKthLargest(nums []int, k int) int {
	exists := map[int]int{}
	max := nums[0]

	for _, num := range nums {
		exists[num]++
		if num > max {
			max = num
		}
	}

	for k > 0 {
		if count, ok := exists[max]; ok {
			k -= count // Decrease k by the count of this number
			if k <= 0 {
				return max
			}
		}
		max-- // Move to the next lower number
	}

	return max
}
