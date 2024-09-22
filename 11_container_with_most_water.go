package main

import "math"

func maxArea(height []int) int {
	max_water := 0

	for left, right := 0, len(height)-1; left < right; {
		lower_height := int(math.Min(float64(height[left]), float64(height[right])))
		water := (right - left) * lower_height
		if water > max_water {
			max_water = water
		}
		if height[left] == lower_height {
			left++
			continue
		}
		right--
	}

	return max_water
}
