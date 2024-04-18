package main

import "math"

func reverse_integer(x int) int {
	if math.Abs(float64(x)) >= 1000000000 && x%10 >= 3 {
		return 0
	}
	res := 0
	for x != 0 {
		res = res*10 + x%10
		if res < -2147483648 || res > 2147483647 {
			return 0
		}
		x /= 10
	}
	return res
}
