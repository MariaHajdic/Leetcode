package main

func numIdenticalPairs(nums []int) int {
	pairs := map[int]int{}

	for _, num := range nums {
		pairs[num] += 1
	}

	result := 0
	for _, val := range pairs {
		if val > 1 {
			result += val * (val - 1) / 2
		}
	}

	return result
}
