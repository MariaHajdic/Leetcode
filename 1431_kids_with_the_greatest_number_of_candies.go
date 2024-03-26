package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := candies[0]
	for _, candy := range candies {
		if candy > max {
			max = candy
		}
	}

	result := make([]bool, len(candies))
	for i, candy := range candies {
		result[i] = candy+extraCandies >= max
	}
	return result
}
