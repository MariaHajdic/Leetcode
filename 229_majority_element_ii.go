package main

func majorityElement2(nums []int) []int {
	counters := map[int]int{}
	for _, num := range nums {
		counters[num]++
	}

	result := []int{}
	bound := len(nums) / 3

	for el := range counters {
		if counters[el] > bound {
			result = append(result, el)
		}
	}

	return result
}
