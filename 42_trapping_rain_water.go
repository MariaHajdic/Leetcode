package main

func trap(height []int) int {
	volume := 0

	left, left_idx := 0, 0
	for i := 0; i < len(height); i++ {
		if left == 0 && height[i] > 0 {
			left = height[i]
			left_idx = i
		} else if height[i] >= left {
			volume += current_volume(height, left_idx, i, left)
			left = height[i]
			left_idx = i
		}
	}

	// The last peak may not be the highest
	right_max := 0
	for i := len(height) - 1; i > left_idx; i-- {
		if height[i] >= right_max {
			right_max = height[i]
		} else {
			volume += right_max - height[i]
		}
	}

	return volume
}

func current_volume(height []int, left, right, max int) int {
	res := 0
	for i := left + 1; i < right; i++ {
		res += max - height[i]
	}
	return res
}
