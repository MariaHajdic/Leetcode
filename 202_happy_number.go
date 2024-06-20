package main

import "math"

func isHappy(n int) bool {
	numbers := map[int]bool{}
	for n != 1 {
		if numbers[n] == true {
			return false
		}
		numbers[n] = true

		new_n := 0
		for temp_n := n; temp_n != 0; temp_n /= 10 {
			new_n += int((math.Pow((float64(temp_n % 10)), 2)))
		}
		n = new_n
	}
	return true
}
